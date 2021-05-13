import plotly.express as px
import sys
import plotly.graph_objects as go
from pathlib import Path


def generatePlot3(fileAbs):
    current = Path.cwd()
    file = current.joinpath(fileAbs)
    xVal = list()
    yVal = list()
    zVal = list()

    with open(file, 'r') as f:
        lines = f.readlines()
        for line in lines:
            line = line.rstrip()
            splits = line.split(',')
            if len(splits) == 3:
                xVal.append(float(splits[0]))
                yVal.append(float(splits[1]))
                zVal.append(float(splits[2]))

    fig = px.scatter_3d(x=xVal, y=yVal, z=zVal)
    fig.write_image(str(Path.joinpath(file.parent, file.stem + ".PNG")))


def generatePlot(fileAbs):
    current = Path.cwd()
    file = current.joinpath(fileAbs)
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
    fig.write_image(str(Path.joinpath(file.parent, file.stem + ".PNG")))


def generateCombinePlot(fileAbsOrig, fileAbsNew):
    current = Path.cwd()
    fileOrig = current.joinpath(fileAbsOrig)
    fileNew = current.joinpath(fileAbsNew)

    xValOrig = list()
    yValOrig = list()

    xValNew = list()
    yValNew = list()

    with open(fileOrig, 'r') as f:
        lines = f.readlines()
        for line in lines:
            line = line.rstrip()
            splits = line.split(',')
            if len(splits) == 2:
                xValOrig.append(float(splits[0]))
                yValOrig.append(float(splits[1]))

    with open(fileNew, 'r') as f:
        lines = f.readlines()
        for line in lines:
            line = line.rstrip()
            splits = line.split(',')
            if len(splits) == 2:
                xValNew.append(float(splits[0]))
                yValNew.append(float(splits[1]))

    fig = go.Figure()
    fig.add_trace(go.Scatter(x=xValOrig,
                             y=yValOrig,
                             mode='markers'))
    fig.add_trace(go.Scatter(x=xValNew,
                             y=yValNew,
                             mode='lines+markers'))
    fig.write_image(str(Path.joinpath(fileNew.parent, fileNew.stem + ".PNG")))


def main():
    if str(sys.argv[1]) == '3':
        generatePlot3(sys.argv[2])
    elif str(sys.argv[1]) == '4':
        generateCombinePlot(sys.argv[2], sys.argv[3])
    else:
        generatePlot(sys.argv[2])


if __name__ == '__main__':
    main()
