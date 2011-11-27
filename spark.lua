
bars = {[0] = "__", "▁ " , "▂ " , "▃ " , "▄ " , "▅ ", "▆ ", "▇ ", "██ "}
vals = {}

if #arg == 0 then
    i = 1
    for line in io.lines() do
        for n in string.gmatch(line, "[%d.]+") do
            vals[i] = n 
            i = i + 1
        end
    end
else
    vals = arg
end

max = 0
min = 3e38 -- max of a float32

for i=1,#vals do
    vals[i] = tonumber(vals[i])
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
