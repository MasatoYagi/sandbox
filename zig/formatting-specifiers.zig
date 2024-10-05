const std = @import("std");

pub fn main() void {
    // 16進数
    // std.fmt.fmtSliceHexLower and std.fmt.fmtSliceHexUpper provide hex formatting for strings as well as {x} and {X} for ints.
    std.debug.print("value: {x}\n", .{123}); // 7b
    std.debug.print("value: {X}\n", .{123}); // 7B
    //// (↑と同じ)
    std.debug.print("value: {}\n", .{std.fmt.fmtSliceHexLower("Zig!")}); // 5a696721
    std.debug.print("value: {}\n", .{std.fmt.fmtSliceHexLower("Zig!")}); // 5a696721

    // 10進数
    // {d} performs decimal formatting for numeric types.
    std.debug.print("value: {d}\n", .{123.4}); // 123.4

    // ASCII
    // {c} formats a byte into an ascii character.
    std.debug.print("value: {c}\n", .{66}); // {

    // メモリサイズ
    // std.fmt.fmtIntSizeDec and std.fmt.fmtIntSizeBin output memory sizes in metric (1000) and power-of-two (1024) based notation.
    std.debug.print("value: {}\n", .{std.fmt.fmtIntSizeDec(1)}); //1B
    std.debug.print("value: {}\n", .{std.fmt.fmtIntSizeBin(1)}); // 1B
    std.debug.print("value: {}\n", .{std.fmt.fmtIntSizeDec(1024)}); // 1.024kB
    std.debug.print("value: {}\n", .{std.fmt.fmtIntSizeBin(1024)}); // 1KiB
    std.debug.print("value: {}\n", .{std.fmt.fmtIntSizeDec(1024 * 1024 * 1024)}); // 1.073741824GB
    std.debug.print("value: {}\n", .{std.fmt.fmtIntSizeBin(1024 * 1024 * 1024)}); // 1GiB

    // 2進数と8進数
    // {b} and {o} output integers in binary and octal format.
    std.debug.print("value: {b}\n", .{254}); // 11111110
    std.debug.print("value: {o}\n", .{254}); // 376

    // ポインタ
    // {*} performs pointer formatting, printing the address rather than the value.
    std.debug.print("value: {*}\n", .{@as(*u8, @ptrFromInt(0xDEADBEEF))}); // u8@deadbeef

    //
    // {e} outputs floats in scientific notation.
    std.debug.print("value: {e}\n", .{3.141592}); // 3.141592e0

    // 文字列
    const hello: [*:0]const u8 = "hello!";
    std.debug.print("value: {s}\n", .{hello}); // hello!
}
