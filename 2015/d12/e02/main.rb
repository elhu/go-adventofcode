require 'json'

def get_sum(data)
  if data.kind_of?(Hash) && data.values.include?("red")
    return 0
  end
  if data.kind_of?(Integer)
    return data
  end
  if data.kind_of?(Array)
    return data.inject(0) { |memo, d| memo + get_sum(d) }
  end
  if data.kind_of?(Hash)
    return get_sum(data.values)
  end
  return 0
end

data = JSON.load(File.read(ARGV[0]))
puts get_sum(data)
