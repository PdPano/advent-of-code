package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	file, _ := os.ReadFile("input1.txt")
	rows := strings.Split(strings.TrimSpace(string(file)), "\n")
    part1(rows)
    part2(rows)
}


func part1(starMap []string){
    expandedMapY := expandMapRows(starMap)
    expandedMap := expandMapRows(transposeMap(expandedMapY))
    coords:=getCoords(expandedMap)

    acc:=0
    for i:=0;i<len(coords);i++{
        for j:=i+1;j<len(coords);j++{
            acc+=absInt(coords[i].i-coords[j].i) + absInt(coords[i].j-coords[j].j)
        }
    }
    fmt.Println(acc)
}

func absInt(v int) int {
    if (v<0){return -v}
    return v
}

func expandMapRows(starMap []string) []string {
    acc:=make([]string,0)
    for i:=0;i<len(starMap);i++{
        currRow := starMap[len(starMap)-1-i]
        if(allDots(currRow)){
            acc = append(acc, currRow)
        }
        acc = append(acc, currRow)
    }
    return acc
}

func allDots(r string) bool{
    return len( strings.ReplaceAll(r,".",""))==0
}

func transposeMap(starMap []string) []string {
    nRows:=len(starMap[0])
    nCols:=len(starMap)
    transposedMap := make([]string, len(starMap[0]))
    for i:=0;i<nRows;i++{
        tmp:=make([]string,0)
        for j:=0;j<nCols;j++{
            tmp = append(tmp, string(starMap[j][i]))
        }
        transposedMap[i]=strings.Join(tmp,"")
    }
    return transposedMap
}

type Coord struct {
    i int
    j int
}

func getCoords(starMap []string)[]Coord{
    acc := make([]Coord,0)
    for i:=0;i<len(starMap);i++{
        for j:=0;j<len(starMap[0]);j++{
            if(starMap[i][j]=='#'){
                acc = append(acc, Coord{i:i,j:j})
            }
        }
    }
    return acc
}

func part2(starMap []string){
    coords:=getCoords(starMap)
    rowIsEmpty := getEmptyRows(starMap)
    colIsEmpty := getEmptyRows(transposeMap(starMap))

    acc:=0
    for i:=0;i<len(coords);i++{
        for j:=i+1;j<len(coords);j++{
            acc+=distDir(coords[i].i, coords[j].i,rowIsEmpty)+distDir(coords[i].j,coords[j].j,colIsEmpty)
        }
    }
    fmt.Println("Part2:",acc)
}

func distDir(i1,i2 int,expanded []bool) int{
    if(i1>i2){
        i1,i2=i2,i1
    }
    acc:=0
    for i:=i1+1;i<=i2;i++{
        if(expanded[i]){
            acc+=1000000
        } else{
            acc++
        }
    }
    return acc
}

func getEmptyRows(starMap []string) []bool{
    rowIsEmpty := make([]bool,len(starMap))
    for i:=0;i<len(starMap);i++{
        rowIsEmpty[i]=allDots(starMap[i])
    }
    return rowIsEmpty
}
