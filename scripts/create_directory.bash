#! /bin/bash

mkdir src/$1
touch src/$1/$1.h src/$1/$1.cpp

echo "#include <gtest/gtest.h>" >> src/$1/$1_test.cpp
echo "" >> src/$1/$1_test.cpp
echo "#include \"$1.h\"" >> src/$1/$1_test.cpp
echo "" >> src/$1/$1_test.cpp
echo "TEST($1, ) {" >> src/$1/$1_test.cpp
echo "}" >> src/$1/$1_test.cpp

echo "#include \"$1.h\"" >> src/$1/main.cpp
echo "" >> src/$1/main.cpp
echo "int main() {" >> src/$1/main.cpp
echo "" >> src/$1/main.cpp
echo "}" >> src/$1/main.cpp
