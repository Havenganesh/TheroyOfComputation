/* construct a dfa that accepts set of all strings over {0,1} of length 2 */

use std::collections::HashMap;

struct DFA {
    initial_state: String,
    accept_states: Vec<String>,
    transition: HashMap<(String, char), String>,
    current_state: String,
}

impl DFA {
    fn new(initial_state: &str, accept_states: Vec<&str>, transition: HashMap<(String, char), String>) -> Self {
        Self {
            initial_state: initial_state.to_string(),
            accept_states: accept_states.into_iter().map(|s| s.to_string()).collect(),
            transition,
            current_state: initial_state.to_string(),
        }
    }

    fn reset(&mut self) {
        self.current_state = self.initial_state.clone();
    }

    fn process(&mut self, symbol: char) {
        if let Some(next_state) = self.transition.get(&(self.current_state.clone(), symbol)) {
            self.current_state = next_state.clone();
        } else {
            // If no valid transition, move to an error state
            self.current_state = "ERROR".to_string();
        }
    }

    fn accept(&self) -> bool {
        self.accept_states.contains(&self.current_state)
    }

    fn process_string(&mut self, string: &str) -> bool {
        self.reset();
        for symbol in string.chars() {
            self.process(symbol);
            if self.current_state == "ERROR" {
                return false;
            }
        }
        self.accept()
    }
}

fn main() {
    let mut transition: HashMap<(String, char), String> = HashMap::new();
    transition.insert(("q0".to_string(), '0'), "q1".to_string());
    transition.insert(("q0".to_string(), '1'), "q1".to_string());
    transition.insert(("q1".to_string(), '0'), "q2".to_string());
    transition.insert(("q1".to_string(), '1'), "q2".to_string());

    let accept_states = vec!["q2"];
    let mut dfa = DFA::new("q0", accept_states, transition);

    let test_strings = vec![
        "0", "1", "00", "01", "10", "11", "000", "001", "010", "011", "100", "101", "110", "111"
    ];

    for s in test_strings {
        println!("String: {}, Accepted: {}", s, dfa.process_string(s));
    }
}
