data = File.read(ARGV[0])
gates = data.split("\n\n")[1]

gates = gates.split("\n").flat_map do |gate|
  gate = gate.split(" ")
  [[gate[0], gate[4]], [gate[2], gate[4]]]
end

File.open("dot.dot", "w+") do |fh|
  fh.puts "digraph G {"
  gates.each do |gate|
    fh.puts "  #{gate[0]} -> #{gate[1]};"
  end
  fh.puts "}"
end
