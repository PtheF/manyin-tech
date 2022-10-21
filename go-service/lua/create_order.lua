
--
---- return value:
---- 1 order create ok
---- 2 stock not enough
---- 3 order already exist
----
---- order status:
---- 0 not pay
---- 1 payed
--


-- 业务前缀 {order}
local key_pre = KEYS[1]

-- 购买数量
local prod_num = ARGV[1]

-- 购买商品 id
local prod_id = ARGV[2]

-- 订单 id
local order_id = ARGV[3]

-- 用于检测订单的令牌
local order_token = ARGV[4]

-- 订单所属用户
local user_id = ARGV[5]

-- order-info
local order_info = ARGV[6]



-- token key {order}:create:token:sdkfjidsli14324
local order_token_key = '{order}:create:token:' .. order_token

-- product stock key {order}:item:stock:prod_id
local item_stock_key = key_pre .. ':item:stock:' .. prod_id

-- order key = {order}:id:order_id
local order_key = key_pre .. ':id:' .. order_id

-- user orders
local user_order_key = key_pre .. ":user:" .. user_id



-- 1. verify orderToken exist
if not redis.call('get', order_token_key) then

    -- order exist
    return 3
end


-- 2. verify stock enough
-- if ( tonumber( redis.call('get', item_stock_key) ) - tonumber(prod_num)  < 0 ) then
--     return 2
-- end

-- verify ok, stock enough, update stock
-- redis.call('incrby', item_stock_key, -1 * tonumber(prod_num))

-- set order status 0 = not pay
redis.call('hset', order_key, 'status', 0)

-- set order info
redis.call('hset', order_key, 'info', order_info)

-- set owner
redis.call("hset", order_key, "user", user_id)

-- add user order
redis.call('sadd', user_order_key, order_id)

-- del order token
redis.call('del', order_token_key)

-- set expire, order should be payed in 30 min
redis.call("expire", order_key, 1860)

return 1

