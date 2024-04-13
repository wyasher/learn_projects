cmake_minimum_required(VERSION 3.25)
project(simu8051 C)

set(CMAKE_C_STANDARD 11)
# 加入hex
file(GLOB HEX_FILES "hex/*.hex")
foreach (HEX_FILE ${HEX_FILES})
    configure_file(${HEX_FILE} ${CMAKE_CURRENT_BINARY_DIR}/${DATA_FILE} COPYONLY)
endforeach ()
add_executable(simu8051 main.c simu8051.h simu8051.c simu8051_tools.h simu8051_tools.c)
