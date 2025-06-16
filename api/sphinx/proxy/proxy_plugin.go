package api

import "github.com/NpoolPlatform/kunman/message/sphinx/proxy"

func (s *Server) ProxyPlugin(stream sphinxproxy.SphinxProxy_ProxyPluginServer) error {
	newPluginStream(stream)
	return nil
}
