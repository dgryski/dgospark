bars = {[0] = "__", "▁ " , "▂ " , "▃ " , "▄ " , "▅ ", "▆ ", "▇ ", "█ "}

max = 0
min = 3e38 -- max of a float32

vals = {}

for i=1,#arg do
    vals[i]=arg[i] + 0 -- coerce to number
    if max < vals[i] then max = vals[i] end
    if min > vals[i] then min = vals[i] end
end

scale=(#bars)/ (max-min)
graph = ""
for i=1,#vals do
    h = math.floor(scale * (vals[i]-min))
    graph = graph .. bars[h]
end

print ("|" .. graph .. "|")
