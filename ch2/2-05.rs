
type BytePointer = *const std::vec::Vec<char>;

// This function should print std::ascii
fn show_bytes(start: BytePointer, length: u32) {
    println!("{:?}", start);
    println!("{:?}", length);

    let chars = "abc";

    for c in chars.chars() { 
        println!("{:02x} ", c.to_digit(16))
    }


    for i in 0..length {
        print!("{:02x} ", i);
        
    }
    println!(assert_eq!('z'.to_digit(16), None))
    
}

fn main() {
    // Convert string into character vector
    let s = "Hello world!";
    let char_vec: Vec<char> = s.chars().collect();

    // Print character vector
    println!("{:?}", char_vec);

    // Assign char vector
    let x: BytePointer = &char_vec;

    // Print pointer
    println!("{:?}", x);

    show_bytes(x, 10)
}
