# encoding: utf-8
# copyright: 2018, Chef Software, Inc.
# license: All rights reserved

title 'IAM v2 migration API integration tests'

control 'upgrade-iam-v2-1' do
  title 'IAM v2 Policies'
  desc 'test IAM v2 policy API post migration'

  describe 'list policies' do
    it 'includes the IAM v2 default policies' do
      resp = automate_api_request('/apis/iam/v2/policies')
      expect(resp.http_status).to eq 200

      all_policies = resp.parsed_response_body[:policies]
      other_policies = all_policies
        .reject{ |p| /^[a-f0-9-]+ \(custom\)$/.match(p[:name]) }
        .reject{ |p| /^\[Legacy\]/.match(p[:name]) }

      # default policies automatically added during migration
      expect(other_policies.length).to be >= 4
    end

    it 'the editors default policy includes editor role' do
      resp = automate_api_request('/apis/iam/v2/policies')
      expect(resp.http_status).to eq 200

      all_policies = resp.parsed_response_body[:policies]
      policies = all_policies.select{ |p| /^team:local:viewers$/.match(p[:members][0]) }
      expect(policies.length).to eq 1
    end

    it 'the viewers default policy includes editor role' do
      resp = automate_api_request('/apis/iam/v2/policies')
      expect(resp.http_status).to eq 200

      all_policies = resp.parsed_response_body[:policies]
      policies = all_policies.select{ |p| /^team:local:editors$/.match(p[:members][0]) }
      expect(policies.length).to eq 1
    end
  end
end

control 'upgrade-iam-v2-3' do
  title 'IAM v2 admin token'
  desc 'test IAM v2 admin token post migration'

  # only run this if we've been handed an admin token to use here
  only_if { ENV['INSPEC_UPGRADE_IAM_V2_3_ADMIN_TOKEN'] }

  describe 'IAM v2 admin token' do
    it 'GET iam/v2/policies returns 200 for admin token generated by iam token create --admin' do
      token = ENV['INSPEC_UPGRADE_IAM_V2_3_ADMIN_TOKEN']

      expect(
        automate_client_api_request(
          '/apis/iam/v2/policies',
          token,
          http_method: 'GET',
        ).http_status
      ).to eq(200)
    end
  end
end
