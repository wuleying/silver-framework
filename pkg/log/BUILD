package(licenses = ["notice"])  # Apache 2.0

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "log",
    srcs = [
        "log.go",
    ],
    importpath = "github.com/wuleying/silver-framework/pkg/log",
    visibility = [
        "//visibility:public",
    ],
    deps = ["//pkg/linewriter"],
)