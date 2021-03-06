package redisc

import "fmt"
import "strconv"
import "github.com/garyburd/redigo/redis"

func Build_hash_key(project, dimension, key, calculation, interval string) string {
	values := []interface{}{"hash:", project, ":", dimension, ":", key, ":", calculation, ":", interval}
	hashkey := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", values...)
	return hashkey
}

func Build_primary_key(project, dimension, key, primarykey string) string {
	values := []interface{}{project, ":", dimension, ":", key, ":", primarykey}
	myprimarykey := fmt.Sprintf("%s%s%s%s%s%s%s", values...)
	return myprimarykey
}

func Build_set_key(project, dimension, key string) string {
	values := []interface{}{project, ":", dimension, ":", key, ":set"}
	setkey := fmt.Sprintf("%s%s%s%s%s%s", values...)
	return setkey
}

func Get_calculated_data(dbnumber, project, dimension, key, calculation, interval string) string {
	cfg := NewRedisConfig()
	connect_string := cfg.Connect_string()
	c, err := redis.Dial("tcp", connect_string)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	redis.String(c.Do("SELECT", dbnumber))
	hashkey := Build_hash_key(project, dimension, key, calculation, interval)
	//fmt.Println(dbnumber, " ", hashkey)
	hstrings, err := redis.Strings(c.Do("HGETALL", hashkey))

	if err != nil {
		panic(err)
	}

	sa := make([]string, 0)
	sa, err = getpairs(sa, hstrings...)
	result := fmt.Sprintf("%s", sa)
	return result
}

func Get_event_data(dbnumber, project, dimension, key string) string {
	cfg := NewRedisConfig()
	connect_string := cfg.Connect_string()
	c, err := redis.Dial("tcp", connect_string)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	redis.String(c.Do("SELECT", dbnumber))
	setkey := Build_set_key(project, dimension, key)

	primarykeys, err := redis.Strings(c.Do("SMEMBERS", setkey))

	if err != nil {
		fmt.Println(err)
		return ("Get_event_data redis error getting primarykeys")
	}

	if len(primarykeys) < 1 {
		return "No primary keys"
	}

	//fmt.Printf("%v\n", primarykeys)

	sa := make([]string, 0)

	for pk := range primarykeys {
		//fmt.Println(pk)
		pkstr := strconv.Itoa(pk)
		hashkey := Build_primary_key(project, dimension, key, pkstr)
		hstrings, err := redis.Strings(c.Do("HGETALL", hashkey))

		if err != nil {
			fmt.Println(err)
			return ("Get_event_data redis error getting hashkey")
		}

		//fmt.Println(hashkey)
		//fmt.Println(hstrings)
		sa, err = getpairs(sa, hstrings...)
	}

	result := fmt.Sprintf("%s", sa)
	return result
}
