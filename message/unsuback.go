// Copyright (c) 2014 The SurgeMQ Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package message

// UnSubAckMessage The UNSUBACK Packet is sent by the Server to the Client to confirm receipt of an
// UNSUBSCRIBE Packet.
type UnSubAckMessage struct {
	PubAckMessage
}

var _ Message = (*UnSubAckMessage)(nil)

// NewUnSubAckMessage creates a new UNSUBACK message.
func NewUnSubAckMessage() *UnSubAckMessage {
	msg := &UnSubAckMessage{}
	msg.SetType(UNSUBACK) // nolint: errcheck

	return msg
}
