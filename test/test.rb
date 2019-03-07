(0..3).each do |c|
  `ab -n1000 http://localhost:300#{c}/ > #{c}`
end
