package db

import (
	"github.com/garyburd/redigo/redis" // Não sei se esse repositório para trabalhar com Redis é o melhor, TODO: Procurar algo melhor
)

const (
	redisExpire = 60 // Guarda o tempo de expiração do Redis DÚVIDA: esse tempo é em minutos ou segundos?
)

func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", "redis:6379") // Tipo de Conexão, endereço do Host e porta de conexão
	HandleError(err)
	return c
}

func Set(key string, value []byte) error {

	conn := RedisConnect()
	defer conn.Close()

	_, err := conn.Do("SET", key, []byte(value))
	HandleError(err)

	conn.Do("EXPIRE", key, redisExpire)

	return err
}

func Get(key string) ([]byte, error) {

	conn := RedisConnect()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	HandleError(err)

	return data, err
}

// Flush é responsável por excluir uma key quando o cache for atualizado e uma nova chave for inserida
func Flush(key string) ([]byte, error) {

	conn := RedisConnect()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("DEL", key))
	HandleError(err)

	return data, err
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
