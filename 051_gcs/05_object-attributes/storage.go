package skyhdd

import (
	"golang.org/x/net/context"
	"google.golang.org/cloud/storage"
	"io"
)

func putFile(ctx context.Context, name string, rdr io.Reader) error {

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	writer := client.Bucket(gcsBucket).Object(name).NewWriter(ctx)
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	io.Copy(writer, rdr)
	return writer.Close()
}

func getAttrs(ctx context.Context, name string) (*storage.ObjectAttrs, error) {

	var attributes *storage.ObjectAttrs
	client, err := storage.NewClient(ctx)
	if err != nil {
		return attributes, err
	}
	defer client.Close()

	attributes, err = client.Bucket(gcsBucket).Object(name).Attrs(ctx)
	if err != nil {
		return attributes, err
	}
	return attributes, nil
}