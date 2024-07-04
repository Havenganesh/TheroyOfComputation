class DFA:
    def __init__(self, initial_state, accept_states, transition):
        self.initial_state = initial_state
        self.accept_states = set(accept_states)
        self.transition = transition
        self.current_state = initial_state

    def reset(self):
        self.current_state = self.initial_state

    def process(self, symbol):
        if symbol in self.transition[self.current_state]:
            self.current_state = self.transition[self.current_state][symbol]
        else:
            # If no valid transition, move to an error state
            self.current_state = None

    def accept(self):
        return self.current_state in self.accept_states

    def process_string(self, string):
        self.reset()
        for symbol in string:
            self.process(symbol)
            if self.current_state is None:
                return False
        return self.accept()


# Define transitions for the DFA
transition = {
    'q0': {'0': 'q1', '1': 'q0'},
    'q1': {'0': 'q0', '1': 'q1'},
}

# Create a DFA with initial state 'q0' and accept state 'q0'
dfa = DFA('q0', ['q0'], transition)

# Test strings
test_strings = [
    "0",
    "1",
    "00",
    "01",
    "10",
    "11",
    "000",
    "001",
    "010",
    "011",
    "100",
    "101",
    "110",
    "111",
]

for s in test_strings:
    print(f"String: {s}, Accepted: {dfa.process_string(s)}")
