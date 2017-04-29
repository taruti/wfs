package wfs

import (
	"io"
	"net/url"
)

func (s *Server) GetFeatureByTypeTo(w io.Writer, typename string) (int64, error) {
	v := url.Values{}
	v.Add(`service`, `wfs`)
	v.Add(`version`, `1.1.0`)
	v.Add(`request`, `GetFeature`)
	v.Add(`typeName`, typename)
	v.Add(`resultType`, `results`)
	v.Add(`outputFormat`, `gml32`)
	res, err := s.Client.Get(s.Url + `?` + v.Encode())
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	return io.Copy(w, res.Body)
}
