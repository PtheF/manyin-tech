--
---- return string:
---- nil => order paid or expired
---- order_info_json
--


local key_pre = KEYS[1]

local order_id = ARGV[1]

local order_info_key = key_pre .. ":id:" .. order_id


if (redis.call('hget', order_info_key, 'status') ~= '0') then
    return 'nil'
end

redis.call("hset", order_info_key, 'status', '1')

local user_id = redis.call('hget', order_info_key, 'user')
local user_orders_key = key_pre .. ':user:' .. user_id

redis.call('srem', user_orders_key, order_id)


return redis.call('hget', order_info_key, 'info')

