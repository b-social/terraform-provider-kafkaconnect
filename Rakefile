require 'rake'
require 'git'
require 'semantic'

require_relative 'lib/version'

def repo
  Git.open('.')
end

def latest_tag
  repo.tags.map do |tag|
    Semantic::Version.new(tag.name)
  end.max
end

def package_list
  (`go list ./...`).split("\n").join(' ')
end

def source_files
  (`find . -name '*.go' | grep -v vendor`).split("\n")
end

def acceptance_tests
  '.*'
end

def coverage_profile
  'c.out'
end

namespace :provider do
  desc 'Formats all go files'
  task :format do
    source_files.each do |source_file|
      sh "go fmt #{source_file}"
    end
  end

  desc 'Builds the provider'
  task :build do
    sh "go build #{package_list}"
  end

  desc 'Installs the provider'
  task :install => [:build] do
    sh 'go install'
  end

  namespace :test do
    desc 'Runs all unit tests'
    task :unit do
      sh "go test #{package_list}"
    end

    desc 'Runs all acceptance tests'
    task :acceptance do
      ENV['TF_ACC']='1'
      sh "go test #{package_list} -v " +
             "-coverprofile #{coverage_profile} " +
             "-run ^#{acceptance_tests}$"
    end
  end
end

namespace :version do
  task :bump, [:type] do |_, args|
    next_tag = latest_tag.send("#{args.type}!")
    repo.add_tag(next_tag.to_s)
    repo.push('origin', 'master', tags: true)
  end
end

desc 'Creates a release for the current version on Github'
task :release do
  sh 'goreleaser'
end
