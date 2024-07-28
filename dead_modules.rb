class DeadModules < Formula
    desc "CLI tool to manage and delete old node_modules directories"
    homepage "https://github.com/furkando/dead_modules"
    url "https://github.com/furkando/dead_modules/archive/v1.0.0.tar.gz"
    sha256 "d39d54ab503fe617141af3b9068f95dedec269edfc4ba4340f79e491ec4e14cf"
    license "MIT"
  
    depends_on "go" => :build
  
    def install
      system "go", "build", *std_go_args, "-o", bin/"dead_modules"
    end
  
    test do
      assert_match "Dead Modules v1.0.0", shell_output("#{bin}/dead_modules --version")
    end
  end