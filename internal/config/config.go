package config

const RABBITMQ_HOST = "amqp://guest:guest@127.0.0.1:5672/"
const RABBITMQ_CONSUME_QUEUE = "devices"
const REST_API_HOST = "http://localhost:8282/"
const GET_DEVICES = REST_API_HOST + "devices"
const GET_DEVICES_BY_TYPE = REST_API_HOST + "devices/type/%s?page=%d&limit=%d"
const GET_DEVICES_BY_STATUS = REST_API_HOST + "devices/status/%s?page=%d&limit=%d"
const GET_ALL_DEVICES = REST_API_HOST + "devices?page=%d&limit=%d"
