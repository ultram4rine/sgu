use std::sync::{Arc, Mutex};
use std::thread;
use std::time::Instant;

fn main() {
    // Thread-safe counter.
    let counter = Arc::new(Mutex::new(0_i64));
    // Numbers.
    let n: i64 = 100000;
    let nums: Vec<i64> = (0..=n).collect();
    let threads_count = 100;

    let mut handles = vec![];
    let chunks: Vec<Vec<i64>> = nums
        .chunks(nums.clone().len() / threads_count)
        .map(|s| s.to_vec())
        .collect();

    let now = Instant::now();

    for i in 0..threads_count {
        let counter = Arc::clone(&counter);
        let chunk = chunks[i].to_vec();
        let handle = thread::Builder::new()
            .name(format!("thread_{}", i + 1))
            .spawn(move || {
                let thr_cur = thread::current();
                for mut num in chunk {
                    // Count '2'.
                    while num > 0 {
                        if num % 10 == 2 {
                            let mut c = counter.lock().unwrap();
                            *c += 1;
                            println!(
                                "thread {}: {}",
                                thr_cur.name().unwrap_or("unknown thread"),
                                *c
                            );
                        }
                        num /= 10;
                    }
                }
            });
        handles.push(handle.unwrap());
    }

    for handle in handles {
        handle.join().unwrap();
    }

    let elapsed = now.elapsed();
    println!("Result: {}", *counter.lock().unwrap());
    println!("Elapsed: {:.2?}", elapsed);
}
