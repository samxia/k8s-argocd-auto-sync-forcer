package argo

import (
	"os"
	"testing"
)

func Test_writeFile(t *testing.T) {
	type args struct {
		content  []byte
		fileName string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Write content to file",
			args: args{
				content:  []byte("Hello, world!"),
				fileName: "test.txt",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeFile(tt.args.content, tt.args.fileName)
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeFile(tt.args.content, tt.args.fileName)
			fileNameWithPath := "/tmp/" + tt.args.fileName
			// Verify that the file was written successfully
			if _, err := os.Stat(fileNameWithPath); os.IsNotExist(err) {
				t.Errorf("writeFile() did not create the file %s", fileNameWithPath)
			}

			// Read the file and check its content
			fileContent, err := os.ReadFile(fileNameWithPath)
			if err != nil {
				t.Errorf("error reading file: %v", err)
			}

			expectedContent := string(tt.args.content)
			if string(fileContent) != expectedContent {
				t.Errorf("file content = %s; want %s", string(fileContent), expectedContent)
			}

			// Clean up: remove the test file
			if err := os.Remove(fileNameWithPath); err != nil {
				t.Errorf("error removing file: %v", err)
			}
		})
	}
}
func Test_runOSFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Valid Command",
			//echo "echo 'hello'" > /tmp/commands.txt
			args:    args{file: "/tmp/commands.txt"}, // Assuming this file contains valid shell commands
			wantErr: false,
			// Update 'want' according to the expected output of the commands in /tmp/commands.txt
			want: "hello\n",
		},
		{
			name:    "Invalid File",
			args:    args{file: "/tmp/nonexistent.txt"}, // Assuming this file does not exist
			wantErr: true,                               // Expecting an error because the file does not exist
			want:    "",                                 // Since an error is expected, the output string should be empty
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runOSFile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("runOSFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("runOSFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeCommandFileByArgs(t *testing.T) {
	type args struct {
		argocdAppName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Valid Arguments",
			args:    args{argocdAppName: "web-prod"},
			wantErr: false,
			want:    "/tmp/commands.txt",
		},

		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("ARGO_SERVER", "argocd.abc.com")
			os.Setenv("ARGO_USER", "admin")
			os.Setenv("ARGO_PASSWORD", "xxxxxxx")
			got, err := makeCommandFileByArgs(tt.args.argocdAppName)
			if (err != nil) != tt.wantErr {
				t.Errorf("makeCommandFileByArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("makeCommandFileByArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
