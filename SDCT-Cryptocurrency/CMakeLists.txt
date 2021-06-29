# Specify the minimum version for CMake
cmake_minimum_required(VERSION 3.2)

# Project's name 
project(SDCT)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11 -pthread -O3")

find_package(OPENSSL REQUIRED)

include_directories(${OPENSSL_INCLUDE_DIR})

add_executable(test_twisted_elgamal test/test_twisted_elgamal.cpp)
target_link_libraries(test_twisted_elgamal ${OPENSSL_LIBRARIES})

add_executable(test_nizk_dlog_equality test/test_nizk_dlog_equality.cpp)
target_link_libraries(test_nizk_dlog_equality ${OPENSSL_LIBRARIES})

add_executable(test_nizk_plaintext_equality test/test_nizk_plaintext_equality.cpp)
target_link_libraries(test_nizk_plaintext_equality ${OPENSSL_LIBRARIES})

add_executable(test_nizk_plaintext_knowledge test/test_nizk_plaintext_knowledge.cpp)
target_link_libraries(test_nizk_plaintext_knowledge ${OPENSSL_LIBRARIES})

add_executable(test_ipproof test/test_ipproof.cpp)
target_link_libraries(test_ipproof ${OPENSSL_LIBRARIES})

add_executable(test_agg_bulletproof test/test_agg_bulletproof.cpp)
target_link_libraries(test_agg_bulletproof ${OPENSSL_LIBRARIES})

add_executable(test_sdct test/test_sdct.cpp)
target_link_libraries(test_sdct ${OPENSSL_LIBRARIES})
