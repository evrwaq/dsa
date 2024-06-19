enable_testing()
find_package(GTest REQUIRED)
include_directories(${GTEST_INCLUDE_DIRS})

file(GLOB_RECURSE TEST_SOURCES "*.cpp")
add_executable(run_tests ${TEST_SOURCES})
target_link_libraries(run_tests ${GTEST_LIBRARIES} pthread dsa_lib)

add_test(NAME run_tests COMMAND run_tests)
