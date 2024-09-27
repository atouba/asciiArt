package main

import (
	"testing"

	"01.gritlab.ax/git/atouba/ascii-art/basic"
)

type testCase struct {
	name     string
	input    string
	expected string
}

var testCases = []testCase{
	{
		name:     "Empty input",
		input:    "",
		expected: "",
	},
	{
		name:     "Single line change",
		input:    "\\n",
		expected: "\n",
	},
	{
		name:  "Hello\\n",
		input: "Hello\\n",
		expected: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                

`,
	},
	{
		name:  "hello",
		input: "hello",
		expected: ` _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
                               
`,
	},

	{
		name:  "HeLlO",
		input: "HeLlO",
		expected: ` _    _          _        _    ____   
| |  | |        | |      | |  / __ \  
| |__| |   ___  | |      | | | |  | | 
|  __  |  / _ \ | |      | | | |  | | 
| |  | | |  __/ | |____  | | | |__| | 
|_|  |_|  \___| |______| |_|  \____/  
                                      
                                      
`,
	},
	{
		name:  "Hello There",
		input: "Hello There",
		expected: ` _    _          _   _                 _______   _                           
| |  | |        | | | |               |__   __| | |                          
| |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___  
|  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \ 
| |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/ 
|_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___| 
                                                                             
                                                                             
`,
	},
	{
		name:  "1Hello 2There",
		input: "1Hello 2There",
		expected: `     _    _          _   _                         _______   _                           
 _  | |  | |        | | | |                ____   |__   __| | |                          
/ | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  
| | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ 
| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ 
|_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| 
                                                                                         
                                                                                         
`,
	},
	{
		name:  "{Hello There}",
		input: "{Hello There}",
		expected: `   __  _    _          _   _                 _______   _                           __    
  / / | |  | |        | | | |               |__   __| | |                          \ \   
 | |  | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___   | |  
/ /   |  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \   \ \ 
\ \   | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/   / / 
 | |  |_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|  | |  
  \_\                                                                              /_/   
                                                                                         
`,
	},
	{
		name:  "Hello\\nThere",
		input: "Hello\\nThere",
		expected: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 
                                       
                                       
`,
	},
	{
		name:  "Hello\\n\\nThere",
		input: "Hello\\n\\nThere",
		expected: ` _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                

 _______   _                           
|__   __| | |                          
   | |    | |__     ___   _ __    ___  
   | |    |  _ \   / _ \ | '__|  / _ \ 
   | |    | | | | |  __/ | |    |  __/ 
   |_|    |_| |_|  \___| |_|     \___| 
                                       
                                       
`,
	},
}

func TestBasic(t *testing.T) {
	for _, tc := range testCases {
		// using t.Run(string, func) produces more details in go test -v
		t.Run(tc.name, func(t *testing.T) {
			result := basic.Basic(tc.input, "", "", "standard", "")
			if tc.expected != result {
				//output = basic.Basic(args[0], "", *clr, "standard", *align)
				t.Errorf("\nInput was \"%s\"\nwant:\n%sgot:\n%s", tc.input, tc.expected, result)
			}
		})
	}
}
