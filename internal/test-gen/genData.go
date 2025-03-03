package test_gen

import (
	"os"
	"strconv"
	"strings"
)

var vars map[string]string

type ConverterFromStr[T any] func(string) T
type GenNum[T any] func(T, T) T
type ConverterToStr[T any] func(T) string

func parseNum[T any](fields []string, to ConverterFromStr[T], from ConverterToStr[T], gen GenNum[T]) string {
	num := gen(to(fields[1]), to(fields[2]))
	s := from(num)
	if len(fields[3]) > 0 {
		vars[fields[3]] = s
	}
	if fields[4] == "false" {
		return ""
	}
	return s + end(fields[5])
}

func unwrap(st, e byte) string {
	var sb strings.Builder
	for i := st; i <= e; i++ {
		sb.WriteByte(i)
	}
	return sb.String()
}

func parseString(fields []string) string {
	var rn strings.Builder
	for i := 0; i < len(fields[1]); i++ {
		if i < len(fields[1])-1 && fields[1][i+1] == '~' {
			rn.WriteString(unwrap(fields[1][i], fields[1][i+2]))
			i += 2
		} else {
			rn.WriteByte(fields[1][i])
		}
	}
	sz, _ := strconv.Atoi(fields[2])
	s := genString(rn.String(), sz)
	if len(fields[3]) > 0 {
		vars[fields[3]] = s
	}
	if fields[4] == "false" {
		return ""
	}
	return s + end(fields[5])
}

func parseVar(fields []string) string {
	return vars[fields[1]] + end(fields[2])
}

func parseFor(lines, fields []string, idx *int) string {
	from, to, it := strToInt(fields[1]), strToInt(fields[2]), strToInt(fields[3])
	var sb strings.Builder
	*idx++
	ix := *idx
	for ; from <= to; from += it {
		for i := ix; lines[i][0] != '}'; i++ {
			*idx = i
			sb.WriteString(parseLine(lines, idx))
		}
	}
	return sb.String() + end(fields[4])
}

func parseLine(lines []string, idx *int) string {
	line := strings.Split(lines[*idx], ";")
	switch line[0] {
	case "integer":
		return parseNum[int64](line, strToInt, intToStr, genInt)
	case "float":
		return parseNum[float64](line, strToFLoat, floatToStr, genFloat)
	case "string":
		return parseString(line)
	case "var":
		return parseVar(line)
	case "for":
		return parseFor(lines, line, idx)
	default:
		return ""
	}
}

func GenData(genScript string) string {
	vars = make(map[string]string)
	file, _ := os.ReadFile(genScript)
	lines := strings.Split(string(file), "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	var res strings.Builder
	for idx := 0; idx < len(lines); idx++ {
		res.WriteString(parseLine(lines, &idx))
	}
	return res.String()
}
