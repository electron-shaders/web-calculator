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
file=open('null.svg',mode='w')
file.write('<?xml version="1.0" standalone="no"?><!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd"><svg t="1632731937314" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1072" data-darkreader-inline-fill="" xmlns:xlink="http://www.w3.org/1999/xlink" width="200" height="200"><defs><style type="text/css">[data-darkreader-inline-fill] { fill: var(--darkreader-inline-fill)  !important; }\n</style></defs><path d="M538.5 850C724.896 850 876 698.896 876 512.5S724.896 175 538.5 175 201 326.104 201 512.5 352.104 850 538.5 850z m0-100C407.332 750 301 643.668 301 512.5S407.332 275 538.5 275 776 381.332 776 512.5 669.668 750 538.5 750z" fill="#979797" p-id="1073" data-darkreader-inline-fill="" style="--darkreader-inline-fill:#535a5d;"></path><path d="M351.578 929.725l475-796c11.32-18.97 5.119-43.526-13.852-54.846-18.97-11.32-43.526-5.12-54.846 13.851l-475 796c-11.32 18.97-5.119 43.526 13.852 54.847 18.97 11.32 43.526 5.118 54.846-13.852z" fill="#333333" p-id="1074" data-darkreader-inline-fill="" style="--darkreader-inline-fill:#262a2b;"></path></svg>')
file.close()