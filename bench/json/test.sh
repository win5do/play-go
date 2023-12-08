#!/usr/bin/env bash
set -euo pipefail

unmarshal_file="output/unmarshal.txt"
unmarshal_diff_file="output/unmarshal-diff.txt"
go test -run='^$' -bench=BenchmarkUnmarshal -count=10 > ${unmarshal_file}
echo -n "" > ${unmarshal_diff_file}
benchstat -col "/S@(std jtr)" -row .name -ignore .file ${unmarshal_file} >> ${unmarshal_diff_file}
echo "---" >> ${unmarshal_diff_file}
benchstat -col "/M@(std jtr)" -row .name -ignore .file ${unmarshal_file} >> ${unmarshal_diff_file}
echo "---" >> ${unmarshal_diff_file}
benchstat -col "/L@(std jtr)" -row .name -ignore .file ${unmarshal_file} >> ${unmarshal_diff_file}

marshal_file="output/marshal.txt"
marshal_diff_file="output/marshal-diff.txt"
go test -run='^$' -bench=BenchmarkMarshal -count=10 > ${marshal_file}
echo  -n "" > ${marshal_diff_file}
benchstat -col "/S@(std jtr)" -row .name -ignore .file ${marshal_file} >> ${marshal_diff_file}
echo "---" >> ${marshal_diff_file}
benchstat -col "/M@(std jtr)" -row .name -ignore .file ${marshal_file} >> ${marshal_diff_file}
echo "---" >> ${marshal_diff_file}
benchstat -col "/L@(std jtr)" -row .name -ignore .file ${marshal_file} >> ${marshal_diff_file}
