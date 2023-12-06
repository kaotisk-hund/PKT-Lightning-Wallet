package votewinnerdb

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/database"
	"github.com/pkt-cash/pktd/pktlog/log"
)

// structure: [blockheight] => [hash][winner]
var bucketName = []byte("votewinnerdb")

func Init(dbTx database.Tx) er.R {
	buck := dbTx.Metadata().Bucket(bucketName)
	if buck == nil {
		log.Infof("Creating vote winner bucket in database")
		if _, err := dbTx.Metadata().CreateBucket(bucketName); err != nil {
			return err
		}
	}
	return nil
}

func Destroy(dbTx database.Tx) er.R {
	buck := dbTx.Metadata().Bucket(bucketName)
	if buck == nil {
		return nil
	}
	log.Infof("Deleting vote winners from database")
	return dbTx.Metadata().DeleteBucket(bucketName)
}

func decodeHeight(h []byte) (int32, er.R) {
	if len(h) != 4 {
		return -1, er.Errorf("Key length is [%d], expecting 4", len(h))
	}
	s := int32(binary.BigEndian.Uint32(h))
	if s < 0 {
		return -1, er.Errorf("Key decodes to [%d] which is less than 0", s)
	}
	return s, nil
}

func bucketAndHeight(dbTx database.Tx, height int32) (database.Bucket, []byte, er.R) {
	buck := dbTx.Metadata().Bucket(bucketName)
	var heightSer [4]byte
	binary.BigEndian.PutUint32(heightSer[:], uint32(height))
	if buck == nil {
		return nil, nil, er.Errorf("Votes not indexed, --votes required")
	}
	return buck, heightSer[:], nil
}

func RemoveWinner(dbTx database.Tx, effectiveHeight int32) er.R {
	if buck, height, err := bucketAndHeight(dbTx, effectiveHeight); err != nil {
		return err
	} else {
		return buck.Delete(height)
	}
}

func PutWinner(dbTx database.Tx, effectiveHeight int32, winner []byte, voteHash []byte) er.R {
	if buck, height, err := bucketAndHeight(dbTx, effectiveHeight); err != nil {
		return err
	} else {
		buf := make([]byte, len(winner)+32)
		if len(voteHash) == 32 {
			copy(buf[:32], voteHash)
		}
		copy(buf[32:], winner)
		log.Debugf("VoteCompute: Storing entry: [%s] => [%s]",
			hex.EncodeToString(height), hex.EncodeToString(buf))
		return buck.Put(height, buf)
	}
}

func ListWinnersBefore(dbTx database.Tx, height int32, handler func(int32, []byte, []byte) er.R) er.R {
	if buck, _, err := bucketAndHeight(dbTx, height); err != nil {
		return err
	} else {
		c := buck.Cursor()
		c.First()
		type ent struct {
			height int32
			val    []byte
		}
		var entries []ent
		for {
			if h, err := decodeHeight(c.Key()); err != nil {
				return err
			} else if h > height {
				break
			} else {
				entries = append(entries, ent{height: h, val: c.Value()})
			}
		}
		log.Debugf("VoteCompute: There exist [%d] winners <= [%d]", len(entries), height)
		for i := len(entries) - 1; i >= 0; i-- {
			val := entries[i].val
			if err := handler(entries[i].height, val[:32], val[32:]); err != nil {
				return err
			}
		}
		return nil
	}
}

func ListWinnersBeforeX(dbTx database.Tx, height int32, handler func(int32, []byte, []byte) er.R) er.R {
	if buck, height, err := bucketAndHeight(dbTx, height); err != nil {
		return err
	} else {
		c := buck.Cursor()
		c.Seek(height)
		for {
			if len(c.Key()) == 0 {
				// Relevant in the first iteration when seek probably did not find the exact entry
			} else if height, err := decodeHeight(c.Key()); err != nil {
				return err
			} else if err := handler(height, c.Value()[:32], c.Value()[32:]); err != nil {
				return err
			}
			if !c.Prev() {
				return nil
			}
		}
	}
}
