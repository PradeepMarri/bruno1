package main

import (
	"github.com/usebruno-api/mcp-server/config"
	"github.com/usebruno-api/mcp-server/models"
	tools_oauth2 "github.com/usebruno-api/mcp-server/tools/oauth2"
	tools_activeglobalenvironmentuid "github.com/usebruno-api/mcp-server/tools/activeglobalenvironmentuid"
	tools_cookies "github.com/usebruno-api/mcp-server/tools/cookies"
	tools_api "github.com/usebruno-api/mcp-server/tools/api"
	tools_custom "github.com/usebruno-api/mcp-server/tools/custom"
	tools_ping "github.com/usebruno-api/mcp-server/tools/ping"
	tools_collectionroot "github.com/usebruno-api/mcp-server/tools/collectionroot"
	tools_redirect_to_ping "github.com/usebruno-api/mcp-server/tools/redirect_to_ping"
	tools_bom_json_test "github.com/usebruno-api/mcp-server/tools/bom_json_test"
	tools_json "github.com/usebruno-api/mcp-server/tools/json"
	tools_cookie "github.com/usebruno-api/mcp-server/tools/cookie"
	tools_bin "github.com/usebruno-api/mcp-server/tools/bin"
	tools_xml_parsed "github.com/usebruno-api/mcp-server/tools/xml_parsed"
	tools_login "github.com/usebruno-api/mcp-server/tools/login"
	tools_request "github.com/usebruno-api/mcp-server/tools/request"
	tools_logout "github.com/usebruno-api/mcp-server/tools/logout"
	tools_text "github.com/usebruno-api/mcp-server/tools/text"
	tools_vars "github.com/usebruno-api/mcp-server/tools/vars"
	tools_basic "github.com/usebruno-api/mcp-server/tools/basic"
	tools_iso_enc "github.com/usebruno-api/mcp-server/tools/iso_enc"
	tools_authorize "github.com/usebruno-api/mcp-server/tools/authorize"
	tools_protected "github.com/usebruno-api/mcp-server/tools/protected"
	tools_callback "github.com/usebruno-api/mcp-server/tools/callback"
	tools_xml_raw "github.com/usebruno-api/mcp-server/tools/xml_raw"
	tools_wsse "github.com/usebruno-api/mcp-server/tools/wsse"
	tools_bearer "github.com/usebruno-api/mcp-server/tools/bearer"
	tools_token "github.com/usebruno-api/mcp-server/tools/token"
	tools_mixed_content_types "github.com/usebruno-api/mcp-server/tools/mixed_content_types"
	tools_resource "github.com/usebruno-api/mcp-server/tools/resource"
	tools_preferences "github.com/usebruno-api/mcp-server/tools/preferences"
	tools_query "github.com/usebruno-api/mcp-server/tools/query"
	tools_headers "github.com/usebruno-api/mcp-server/tools/headers"
	tools_environments "github.com/usebruno-api/mcp-server/tools/environments"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_oauth2.CreatePost_oauth2_authorization_code_tokenTool(cfg),
		tools_activeglobalenvironmentuid.CreateGet_activeglobalenvironmentuidTool(cfg),
		tools_cookies.CreateGet_cookiesTool(cfg),
		tools_api.CreatePost_api_auth_oauth2_authorization_code_tokenTool(cfg),
		tools_custom.CreatePost_customTool(cfg),
		tools_ping.CreateGet_pingTool(cfg),
		tools_api.CreatePost_api_echo_binTool(cfg),
		tools_collectionroot.CreateGet_collectionrootTool(cfg),
		tools_api.CreatePost_api_echo_customTool(cfg),
		tools_api.CreatePost_api_auth_wsse_protectedTool(cfg),
		tools_redirect_to_ping.CreateGet_redirect_to_pingTool(cfg),
		tools_api.CreatePost_api_echo_jsonTool(cfg),
		tools_bom_json_test.CreateGet_bom_json_test_handlerTool(cfg),
		tools_json.CreatePost_jsonTool(cfg),
		tools_cookie.CreateGet_cookie_protectedTool(cfg),
		tools_bin.CreatePost_binTool(cfg),
		tools_api.CreatePost_api_auth_oauth2_client_credentials_tokenTool(cfg),
		tools_json.CreateGet_jsonTool(cfg),
		tools_api.CreateGet_api_auth_oauth2_client_credentials_resourceTool(cfg),
		tools_api.CreateGet_api_auth_cookie_protectedTool(cfg),
		tools_xml_parsed.CreatePost_xml_parsedTool(cfg),
		tools_login.CreatePost_loginTool(cfg),
		tools_api.CreatePost_api_multipart_mixed_content_typesTool(cfg),
		tools_request.CreateGet_requestTool(cfg),
		tools_api.CreatePost_api_auth_oauth2_password_credentials_resourceTool(cfg),
		tools_logout.CreatePost_logoutTool(cfg),
		tools_text.CreatePost_textTool(cfg),
		tools_vars.CreateGet_varsTool(cfg),
		tools_oauth2.CreatePost_oauth2_password_credentials_tokenTool(cfg),
		tools_oauth2.CreatePost_oauth2_authorization_code_resourceTool(cfg),
		tools_api.CreatePost_api_auth_cookie_logoutTool(cfg),
		tools_api.CreatePost_api_auth_basic_protectedTool(cfg),
		tools_basic.CreatePost_basic_protectedTool(cfg),
		tools_oauth2.CreateGet_oauth2_authorization_code_callbackTool(cfg),
		tools_api.CreatePost_api_auth_cookie_loginTool(cfg),
		tools_api.CreateGet_api_auth_oauth2_authorization_code_authorizeTool(cfg),
		tools_api.CreateGet_api_echo_iso_encTool(cfg),
		tools_iso_enc.CreateGet_iso_encTool(cfg),
		tools_api.CreatePost_api_auth_oauth2_authorization_code_resourceTool(cfg),
		tools_authorize.CreateGet_authorizeTool(cfg),
		tools_oauth2.CreateGet_oauth2_client_credentials_resourceTool(cfg),
		tools_protected.CreateGet_protectedTool(cfg),
		tools_protected.CreatePost_protectedTool(cfg),
		tools_callback.CreateGet_callbackTool(cfg),
		tools_api.CreateGet_api_auth_bearer_protectedTool(cfg),
		tools_xml_raw.CreatePost_xml_rawTool(cfg),
		tools_api.CreatePost_api_echo_xml_parsedTool(cfg),
		tools_cookie.CreatePost_cookie_loginTool(cfg),
		tools_wsse.CreatePost_wsse_protectedTool(cfg),
		tools_bearer.CreateGet_bearer_protectedTool(cfg),
		tools_oauth2.CreatePost_oauth2_client_credentials_tokenTool(cfg),
		tools_token.CreatePost_tokenTool(cfg),
		tools_mixed_content_types.CreatePost_mixed_content_typesTool(cfg),
		tools_api.CreateGet_api_echo_bom_json_test_handlerTool(cfg),
		tools_resource.CreateGet_resourceTool(cfg),
		tools_resource.CreatePost_resourceTool(cfg),
		tools_oauth2.CreatePost_oauth2_password_credentials_resourceTool(cfg),
		tools_api.CreateGet_api_auth_oauth2_authorization_code_callbackTool(cfg),
		tools_preferences.CreateGet_preferencesTool(cfg),
		tools_api.CreatePost_api_echo_textTool(cfg),
		tools_oauth2.CreateGet_oauth2_authorization_code_authorizeTool(cfg),
		tools_query.CreateGet_queryTool(cfg),
		tools_api.CreatePost_api_auth_oauth2_password_credentials_tokenTool(cfg),
		tools_headers.CreateGet_headersTool(cfg),
		tools_environments.CreateGet_environmentsTool(cfg),
		tools_cookie.CreatePost_cookie_logoutTool(cfg),
		tools_api.CreatePost_api_echo_xml_rawTool(cfg),
	}
}
