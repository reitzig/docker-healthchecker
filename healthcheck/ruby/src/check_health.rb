require "bundler"
Bundler.require(:default)

target_folder = ENV['SUMMARY_DIR']

summary_json = File.read("#{target_folder}/all.json")
summary = JSON.parse(summary_json, symbolize_names: true)

exit(summary[:all][:status] || 1)
