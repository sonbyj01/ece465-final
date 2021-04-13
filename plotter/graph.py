import plotly.express as px
import pandas as pd

file = '../sample/graph2d-100.txt'
df = pd.DataFrame()
xVal = list()
yVal = list()

with open(file, 'r') as f:
    lines = f.readlines()
    for line in lines:
        line = line.rstrip()
        splits = line.split(',')
        if len(splits) == 2:
            xVal.append(float(splits[0]))
            yVal.append(float(splits[1]))

fig = px.scatter(x=xVal, y=yVal)
fig.write_image('../sample/graph2d-100.png')
