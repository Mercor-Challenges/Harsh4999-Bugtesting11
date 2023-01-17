/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// osintypes holds types that are used to interpret responses from the RangelReale osin server.
// neither osin, nor osincli contain types with the annotations required to have the default json
// marshaller encode and decode them.  Even more unusual: osin does not describe a single type that
// represents the overall return from osin.FinishInfoRequest.  Because of that, a type needs to be
// described in order to make use of the return value, so even if you preferred writing a parser
// you'll end up needing a type anyway.
package osintypes

// InfoResponseData is a type that matches the information returned from osin.FinishInfoRequest (/oauth/info).
type InfoResponseData struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	TokenType        string `json:"token_type"`
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	Expiration       int32  `json:"expires_in"`
}
