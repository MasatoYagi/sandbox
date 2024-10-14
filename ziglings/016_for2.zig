//
// For loops also let you use the "index" of the iteration, a number
// that counts up with each iteration. To access the index of iteration,
// specify a second condition as well as a second capture value.
//
//     for (items, 0..) |item, index| {
//
//         // Do something with item and index
//
//     }
//
// You can name "item" and "index" anything you want. "i" is a popular
// shortening of "index". The item name is often the singular form of
// the items you're looping through.
//
const std = @import("std");

pub fn main() void {
    // Let's store the bits of binary number 1101 in
    // 'little-endian' order (least significant byte or bit first):
    const bits = [_]u8{ 1, 0, 1, 1 };
    var value: u32 = 0;

    // Now we'll convert the binary bits to a number value by adding
    // the value of the place as a power of two for each bit.
    //
    // See if you can figure out the missing pieces:
    for (bits, 2..) |bit, i| {
        // Note that we convert the usize i to a u32 with
        // @intCast(), a builtin function just like @import().
        // We'll learn about these properly in a later exercise.
        const i_u32: u32 = @intCast(i); // pow関数の第3引数に渡すためにu32に変換
        const place_value = std.math.pow(u32, 2, i_u32); // 0桁目は2^0=1, 1桁目は2^1=2, 2桁目は2^2=4, 3桁目は2^3=8（2進数なので2の累乗）
        value += place_value * bit; // 各桁の値 * 2の累乗の値を足し合わて、10進数に変換
    }

    std.debug.print("The value of bits '1101': {}.\n", .{value});
}
//
// As mentioned in the previous exercise, 'for' loops have gained
// additional flexibility since these early exercises were
// written. As we'll see in later exercises, the above syntax for
// capturing the index is part of a more general ability. Hang in
// there!
