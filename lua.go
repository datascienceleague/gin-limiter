package limiter

const Script = `
	local result = {}
		
	local globalKey = KEYS[1] 
	local singleKey = KEYS[2]
	
	local globalLimit = tonumber(ARGV[1])
	local singleLimit = tonumber(ARGV[2])

	local IpGlobalInfo = redis.call('HGETALL', globalKey)
	local IpSingleInfo = redis.call('HGETALL', singleKey)
	
	-- 該Ip第一次訪問
	if #IpGlobalInfo == 0 then
		redis.call('HMSET', globalKey, "Count", 1)
		redis.call('HMSET', singleKey, "Count", 1)
		result[1] = globalLimit - 1 
		result[2] = singleLimit - 1 
		result[3] = "#IpGlobalInfo == 0 Area"
		return result
	end

	-- 從來沒有訪問過這個path
	if #IpSingleInfo == 0 then 
		redis.call('HMSET', singleKey, "Count", 1)
		result[2] = singleLimit - 1
		result[3] = "#IpSingleInfo == 0 Area"
		IpSingleInfo = redis.call('HGETALL', singleKey)
	end

	local globalExpired = ARGV[3]
	local singleExpired = ARGV[4]

	if globalExpired == true then 
		redis.call('HMSET', globalKey, "Count", 1)
		result[1] = globalLimit - 1
		result[3] = "global expired area"
	end

	if singleExpired == true then
		redis.call('HMSET', singleKey, "Count", 1)
		result[2] = singleLimit - 1
		result[3] = "single expired area"
		if globalExpired == "true" then 
			return result
		end
	end

	local globalCount = tonumber(IpGlobalInfo[2]) 
	local singleCount = tonumber(IpSingleInfo[2]) 

	if globalCount < globalLimit then 
		local NewglobalCount = redis.call('HINCRBY', globalKey, "Count", 1)
		result[1] = globalLimit - NewglobalCount
		result[3] = "global Count < limit"
	else
		result[1] = -1
	end
	
	if singleCount < singleLimit then 
		local NewsingleCount = redis.call('HINCRBY', singleKey, "Count", 1)
		result[2] = singleLimit - NewsingleCount
		result[3] = "single Count < limit"
	else 
		result[2] = -1
	end

	return result
`

const TestScript = `
	local result = {}

	-- 測試註解
	result[5] = 28
	result[1] = 10
	result[2] = 20 + 2
	result[3] = KEYS[1]
	result[4] = ARGV[1]
	
	return result
`
