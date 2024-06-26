

local key = KEYS[1]

-- count 使用次数，也就是验证次数
local cntKey = key..":cnt"

-- 用户输入的验证码
local expectedCode = ARGV[1]

local cnt = tonumber(redis.call("get", cntKey))
local code = redis.call("get", key)

-- 用户一直输错 验证次数耗尽了 合理怀疑恶搞
if cnt == nil or cnt <= 0 then
    return -1
end

-- 输入对了
if code == expectedCode then
    -- 可验证次数 归零
    redis.call("set", cntKey, 0)
    return 0

-- 用户手一抖 输错了
else
    -- 可验证次数 -1
    redis.call("decr", cntKey)
    return -2
end