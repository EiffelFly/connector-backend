package resource

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)

// ExtractFromMetadata extracts context metadata given a key
func ExtractFromMetadata(ctx context.Context, key string) ([]string, bool) {
	data, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return []string{}, false
	}
	return data[strings.ToLower(key)], true
}

// GetCollectionID returns the resource collection ID given a resource name
func GetCollectionID(name string) (string, error) {
	colID := name[:strings.LastIndex(name, "/")]
	if colID == "" {
		return "", fmt.Errorf("Error when extract resource collection id from resource name %s", name)
	}
	if strings.LastIndex(colID, "/") != -1 {
		colID = colID[strings.LastIndex(colID, "/")+1:]
	}
	return colID, nil
}

// GetRscNameID returns the resource ID given a resource name
func GetRscNameID(name string) (string, error) {
	id := name[strings.LastIndex(name, "/")+1:]
	if id == "" {
		return "", fmt.Errorf("Error when extract resource id from resource name '%s'", name)
	}
	return id, nil
}

// GetPermalinkUID returns the resource UID given a resource permalink
func GetPermalinkUID(permalink string) (string, error) {
	uid := permalink[strings.LastIndex(permalink, "/")+1:]
	if uid == "" {
		return "", fmt.Errorf("Error when extract resource id from resource permalink '%s'", permalink)
	}
	return uid, nil
}

// GetOwner returns the resource owner
func GetOwner(ctx context.Context) (string, error) {
	metadatas, ok := ExtractFromMetadata(ctx, "owner")
	if ok {
		if len(metadatas) == 0 {
			return "", fmt.Errorf("Cannot find owner in your request")
		}
		return metadatas[0], nil
	}
	return "", fmt.Errorf("Error when extract owner metadata")
}
