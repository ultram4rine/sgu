use rand::Rng;
use std::thread;

fn main() {
    let mut rng = rand::thread_rng();
    let nums: Vec<i32> = (1..=10000000).map(|_| rng.gen_range(1..=1000)).collect();
    let threads_count = 100;
    let chunks: Vec<Vec<i32>> = nums
        .chunks(nums.clone().len() / threads_count)
        .map(|s| s.to_vec())
        .collect();

    let handles = (0..threads_count).into_iter().map(|i| {
        let chunk = chunks[i].to_vec();
        thread::Builder::new()
            .name(format!("thread_{}", i))
            .spawn(move || process_nums(chunk))
            .unwrap()
    });

    handles.into_iter().for_each(|h| h.join().unwrap());
}

fn process_nums(nums: Vec<i32>) {
    let handle = thread::current();
    for num in nums {
        if num % 3 == 0 {
            println!(
                "{}: {} is dividable by 3",
                handle.name().unwrap_or("unknown thread"),
                num
            );
        }
    }
}
