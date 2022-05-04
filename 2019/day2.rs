use std::fs;

fn main() {
    let finput = fs::read_to_string("./input/day2.txt").expect("Something went wrong!");
    let filist = finput.split_terminator(",");

    // Fill out a vector of integers from our array of strings
    let mut inp: Vec<i32> = Vec::new();
    for i in filist {
        let p: i32 = i.parse().unwrap();
        inp.push(p);
    };

    for i in 0..inp.len() {
        println!("{}", handle_opcode(i, &inp));
    };

}

// Takes the current position, the mutable vector of operations,
// performs the asked computions in the challenge and then returns
// the next pos in the ops vector to access for the next run.
// If we return -1 just continue to the next object I guess.
fn handle_opcode(pos: usize, ops: &Vec<i32>) -> i32 {
    if ops[pos] == 1 {
        ops[pos]
    } else if ops[pos] == 2 {
        ops[pos]
    } else {
        -1
    }
}