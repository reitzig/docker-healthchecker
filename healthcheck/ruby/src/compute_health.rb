require "bundler"
Bundler.require(:default)
require_relative 'shared.rb'

checks_interval = ENV['INTERVAL'].to_i
target_folder = ENV['SUMMARY_DIR']

checks_file = ARGV[0]
checks_json = File.read(checks_file)
checks = JSON.parse(checks_json, symbolize_names: true)

while true
    puts "Computing health"

    summary = { all: { status: 0 } }
    all_start_time = Time.now
    checks.each do |check|
        puts " - #{check[:description]}?"

        check_summary = {}
        start_time = Time.now

        # TODO: compute check
        sleep 1

        check_summary[:duration] = pretty_duration(Time.now, start_time)
        summary[check[:description]] = check_summary
    end
    summary[:all][:duration] = pretty_duration(Time.now, all_start_time)
    summary[:all][:finished] = Time.now

    File.open("#{target_folder}/all.json", 'w') { |f| f.write(JSON.pretty_generate(summary)) }
    FileUtils.chmod(0644, "#{target_folder}/all.json")

    sleep checks_interval
end
