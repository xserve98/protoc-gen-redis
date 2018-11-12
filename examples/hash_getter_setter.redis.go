// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hash_getter_setter.proto

package test

import github_com_gomodule_redigo_redis "github.com/gomodule/redigo/redis"
import github_com_mitchellh_mapstructure "github.com/mitchellh/mapstructure"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/galaxyobe/protoc-gen-redis/proto"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// new HashGetterAndSetterType redis controller with redis pool
func (m *HashGetterAndSetterType) RedisController(pool *github_com_gomodule_redigo_redis.Pool) *HashGetterAndSetterTypeRedisController {
	return &HashGetterAndSetterTypeRedisController{
		pool: pool,
		m:    m,
	}
}

// HashGetterAndSetterType redis controller
type HashGetterAndSetterTypeRedisController struct {
	pool *github_com_gomodule_redigo_redis.Pool
	m    *HashGetterAndSetterType
}

// new HashGetterAndSetterType redis controller with redis pool
func NewHashGetterAndSetterTypeRedisController(pool *github_com_gomodule_redigo_redis.Pool) *HashGetterAndSetterTypeRedisController {
	return &HashGetterAndSetterTypeRedisController{pool: pool, m: new(HashGetterAndSetterType)}
}

// get HashGetterAndSetterType
func (r *HashGetterAndSetterTypeRedisController) HashGetterAndSetterType() *HashGetterAndSetterType {
	return r.m
}

// set HashGetterAndSetterType
func (r *HashGetterAndSetterTypeRedisController) SetHashGetterAndSetterType(m *HashGetterAndSetterType) {
	r.m = m
}

// load HashGetterAndSetterType from redis hash
func (r *HashGetterAndSetterTypeRedisController) Load(key string) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// load data from redis hash
	data, err := github_com_gomodule_redigo_redis.ByteSlices(conn.Do("HGETALL", key))
	if err != nil {
		return err
	}

	// parse redis hash field name and value
	structure := make(map[string]interface{})
	for i := 0; i < len(data); i += 2 {
		switch string(data[i]) {
		default:
			structure[string(data[i])] = string(data[i+1])
		}
	}

	// use mapstructure weak decode structure to HashGetterAndSetterType
	return github_com_mitchellh_mapstructure.WeakDecode(structure, r.m)
}

// get HashGetterAndSetterType field from redis hash return string value
func (r *HashGetterAndSetterTypeRedisController) GetString(key string, field string) (value string, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return github_com_gomodule_redigo_redis.String(conn.Do("HGET", key, field))
}

// get HashGetterAndSetterType field from redis hash return bool value
func (r *HashGetterAndSetterTypeRedisController) GetBool(key string, field string) (value bool, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return github_com_gomodule_redigo_redis.Bool(conn.Do("HGET", key, field))
}

// get HashGetterAndSetterType field from redis hash return int64 value
func (r *HashGetterAndSetterTypeRedisController) GetInt64(key string, field string) (value int64, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return github_com_gomodule_redigo_redis.Int64(conn.Do("HGET", key, field))
}

// get HashGetterAndSetterType field from redis hash return uint64 value
func (r *HashGetterAndSetterTypeRedisController) GetUint64(key string, field string) (value uint64, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return github_com_gomodule_redigo_redis.Uint64(conn.Do("HGET", key, field))
}

// get HashGetterAndSetterType field from redis hash return float64 value
func (r *HashGetterAndSetterTypeRedisController) GetFloat64(key string, field string) (value float64, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get field
	return github_com_gomodule_redigo_redis.Float64(conn.Do("HGET", key, field))
}

// store HashGetterAndSetterType to redis hash
func (r *HashGetterAndSetterTypeRedisController) Store(key string) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// make args
	args := make([]interface{}, 0)

	// add redis key
	args = append(args, key)

	// add redis field and value
	args = append(args, "SomeString", r.m.SomeString)
	args = append(args, "SomeBool", r.m.SomeBool)
	args = append(args, "SomeInt32", r.m.SomeInt32)
	args = append(args, "SomeUint32", r.m.SomeUint32)
	args = append(args, "SomeInt64", r.m.SomeInt64)
	args = append(args, "SomeUint64", r.m.SomeUint64)
	args = append(args, "SomeFloat", r.m.SomeFloat)
	args = append(args, "SomeEnum", int32(r.m.SomeEnum))

	// use redis hash store HashGetterAndSetterType data
	_, err := conn.Do("HMSET", args...)

	return err
}

// store HashGetterAndSetterType to redis hash with key and ttl expire second
func (r *HashGetterAndSetterTypeRedisController) StoreWithTTL(key string, ttl uint64) error {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// make args
	args := make([]interface{}, 0)

	// add redis key
	args = append(args, key)

	// add redis field and value
	args = append(args, "SomeString", r.m.SomeString)
	args = append(args, "SomeBool", r.m.SomeBool)
	args = append(args, "SomeInt32", r.m.SomeInt32)
	args = append(args, "SomeUint32", r.m.SomeUint32)
	args = append(args, "SomeInt64", r.m.SomeInt64)
	args = append(args, "SomeUint64", r.m.SomeUint64)
	args = append(args, "SomeFloat", r.m.SomeFloat)
	args = append(args, "SomeEnum", int32(r.m.SomeEnum))

	// use redis hash store HashGetterAndSetterType data with expire second
	err := conn.Send("MULTI")
	if err != nil {
		return err
	}
	err = conn.Send("HMSET", args...)
	if err != nil {
		return err
	}
	err = conn.Send("EXPIRE", key, ttl)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXEC")

	return err
}

// set HashGetterAndSetterType field value to redis hash
func (r *HashGetterAndSetterTypeRedisController) SetFieldValue(key string, field string, value interface{}) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set field
	_, err = conn.Do("HSET", key, field, value)

	return
}

// set HashGetterAndSetterType SomeString field with key and SomeString
func (r *HashGetterAndSetterTypeRedisController) SetSomeString(key string, someString string) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeString field
	r.m.SomeString = someString
	_, err = conn.Do("HSET", key, "SomeString", someString)

	return
}

// set HashGetterAndSetterType SomeBool field with key and SomeBool
func (r *HashGetterAndSetterTypeRedisController) SetSomeBool(key string, someBool bool) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeBool field
	r.m.SomeBool = someBool
	_, err = conn.Do("HSET", key, "SomeBool", someBool)

	return
}

// set HashGetterAndSetterType SomeInt32 field with key and SomeInt32
func (r *HashGetterAndSetterTypeRedisController) SetSomeInt32(key string, someInt32 int32) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeInt32 field
	r.m.SomeInt32 = someInt32
	_, err = conn.Do("HSET", key, "SomeInt32", someInt32)

	return
}

// set HashGetterAndSetterType SomeUint32 field with key and SomeUint32
func (r *HashGetterAndSetterTypeRedisController) SetSomeUint32(key string, someUint32 uint32) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeUint32 field
	r.m.SomeUint32 = someUint32
	_, err = conn.Do("HSET", key, "SomeUint32", someUint32)

	return
}

// set HashGetterAndSetterType SomeInt64 field with key and SomeInt64
func (r *HashGetterAndSetterTypeRedisController) SetSomeInt64(key string, someInt64 int64) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeInt64 field
	r.m.SomeInt64 = someInt64
	_, err = conn.Do("HSET", key, "SomeInt64", someInt64)

	return
}

// set HashGetterAndSetterType SomeUint64 field with key and SomeUint64
func (r *HashGetterAndSetterTypeRedisController) SetSomeUint64(key string, someUint64 uint64) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeUint64 field
	r.m.SomeUint64 = someUint64
	_, err = conn.Do("HSET", key, "SomeUint64", someUint64)

	return
}

// set HashGetterAndSetterType SomeFloat field with key and SomeFloat
func (r *HashGetterAndSetterTypeRedisController) SetSomeFloat(key string, someFloat float32) (err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// set SomeFloat field
	r.m.SomeFloat = someFloat
	_, err = conn.Do("HSET", key, "SomeFloat", someFloat)

	return
}

// get HashGetterAndSetterType SomeEnum field value with key
func (r *HashGetterAndSetterTypeRedisController) GetSomeEnum(key string) (someEnum HashGetterAndSetterType_Enum, err error) {
	// redis conn
	conn := r.pool.Get()
	defer conn.Close()

	// get SomeEnum field
	if value, err := github_com_gomodule_redigo_redis.Int64(conn.Do("HGET", key, "SomeEnum")); err != nil {
		return someEnum, err
	} else {
		r.m.SomeEnum = HashGetterAndSetterType_Enum(value)
	}

	return r.m.SomeEnum, nil
}
