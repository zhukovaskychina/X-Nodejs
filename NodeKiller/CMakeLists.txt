cmake_minimum_required(VERSION 3.12)
project(NodeKiller)

set(CMAKE_CXX_STANDARD 14)


set(Boost_USE_STATIC_LIBS        ON)
set(Boost_USE_MULTITHREADED      ON)
set(BOOST_ROOT "/usr/include/boost")
find_package(Boost COMPONENTS system filesystem program_options)
if (Boost_FOUND)
    include_directories(${Boost_INCLUDE_DIRS})
    MESSAGE( STATUS "Boost_INCLUDE_DIRS = ${Boost_INCLUDE_DIRS}.")
    MESSAGE( STATUS "Boost_LIBRARIES = ${Boost_LIBRARIES}.")
    MESSAGE( STATUS "Boost_LIB_VERSION = ${Boost_LIB_VERSION}.")

    add_executable(NodeKiller main.cpp)

    link_directories(... ${Boost_LIBRARY_DIRS})

    target_link_libraries(NodeKiller ${Boost_LIBRARIES})
else()
    MESSAGE( STATUS "没发现任何BOOST")
endif ()