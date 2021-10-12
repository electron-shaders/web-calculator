import os

with open("index.html", "r", encoding="utf-8") as f1,open("%s.bak" % "index.html", "w", encoding="utf-8") as f2:
    for line in f1:
        if 'href="' in line:
            line = line.replace('href="', 'href=".')
        if 'src="' in line:
            line = line.replace('src="', 'src=".')
        f2.write(line)
os.remove("index.html")
os.rename("%s.bak" % "index.html", "index.html")

with open("./src/utils/axios.js", "r", encoding="utf-8") as f1,open("%s.bak" % "src/utils/axios.js", "w", encoding="utf-8") as f2:
    for line in f1:
        if 'http://localhost:3001' in line:
            line = line.replace('http://localhost:3001', 'https://demo.xn--ftwm9mhr4a.com')
        f2.write(line)
os.remove("src/utils/axios.js")
os.rename("%s.bak" % "src/utils/axios.js", "src/utils/axios.js")