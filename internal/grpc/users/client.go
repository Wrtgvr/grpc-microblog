package usergrpc

import (
	usersv1 "github.com/wrtgvr/grpc-microblog/gen/users/users"
	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client usersv1.UsersClient
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.NewClient(addr)
	if err != nil {
		return nil, err
	}

	cl := usersv1.NewUsersClient(conn)

	return &Client{
		conn:   conn,
		client: cl,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}
