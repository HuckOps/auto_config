package reader

import (
	"auto_config/encoder"
	"auto_config/encoder/json"
	"auto_config/source"
	"errors"
	"github.com/imdario/mergo"
	"time"
)

type jsonReader struct {
	opts Options
	json encoder.Encoder
}

func (j *jsonReader) Merge(changes ...*source.ChangeSet) (*source.ChangeSet, error) {
	merged := map[interface{}]interface{}{}
	for _, change := range changes {
		codec, ok := j.opts.Encoding[change.Format]
		if !ok {
			return nil, errors.New("get encoder failed")
		}

		var data map[string]interface{}
		// 转为结构体
		if err := codec.Decode(change.Data, &data); err != nil {
			return nil, err
		}
		// 结构体合并
		if err := mergo.Map(&merged, data, mergo.WithOverride); err != nil {
			return nil, err
		}
	}

	// 序列化为json
	b, err := j.json.Encode(merged)
	print("ttt")
	//b, err := json2.Marshal(merged)
	if err != nil {
		panic(err)
		return nil, err
	}
	cs := &source.ChangeSet{
		Timestamp: time.Now(),
		Data:      b,
		Format:    j.json.Type(),
	}
	cs.CheckSum = cs.Sum()

	return cs, nil
}

func NewReader(opts ...Option) Reader {
	options := NewOptions(opts...)
	return &jsonReader{
		json: json.NewEncoder(),
		opts: options,
	}
}
