
-- 验证码在 redis 上的 key，也就是 phone_code:login:188xxxxoooo
local key = KEYS[1]
-- 验证次数，我们一个验证码，最多重复三次
local cntKey = key..":cnt"
-- 你准备的存储的验证码 123456
local val = ARGV[1]

-- 验证码的有效时间是十分钟，600 秒
local ttl = tonumber(redis.call("ttl", key))

-- ttl = -1，如果 Redis 中有这个 Key，但是没有过期时间，说明系统异常（同事手贱，误操作，设置key，没给过期时间）
-- ttl = -2，如果 Redis 中没有这个 key，那么就直接发送
-- ttl > 540，如果 key 有过期时间，但是过期时间还有 9 分钟，发送太频繁，拒绝
-- 否则，重新发送一个验证码
if ttl == -1 then
    -- 系统异常
    return -2
elseif ttl == -2 or ttl < 540 then
    redis.call("set", key, val)
    redis.call("expire", key, 600)
    redis.call("set", cntKey, 3)
    redis.call("expire", cntKey, 600)
    -- 正常
    return 0
else
    -- 发送太频繁
    return -1
end