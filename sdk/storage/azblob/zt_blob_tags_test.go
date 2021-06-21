// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azblob

//
//import (
//	"bytes"
//	"crypto/md5"
//	"encoding/base64"
//	"encoding/binary"
//	"fmt"
//	chk "gopkg.in/check.v1"
//	"io/ioutil"
//	"strings"
//	"time"
//)
//
//func (s *azblobTestSuite) TestSetBlobTags() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//
//	bbClient, _ := getBlockBlobClient(c, containerClient)
//	blobTagsMap := map[string]string{
//		"azure":    "bbClient",
//		"bbClient": "sdk",
//		"sdk":      "go",
//	}
//
//	contentSize := 4 * 1024 * 1024 // 4MB
//	r, _ := getRandomDataAndReader(contentSize)
//
//	blockBlobUploadResp, err := bbClient.Upload(ctx, r, nil)
//	_assert.Nil(err)
//	c.Assert(blockBlobUploadResp.RawResponse.StatusCode, chk.Equals, 201)
//
//	setTagsBlobOptions := SetTagsBlobOptions{
//		BlobTagsMap: &blobTagsMap,
//	}
//	blobSetTagsResponse, err := bbClient.SetTags(ctx, &setTagsBlobOptions)
//	_assert.Nil(err)
//	c.Assert(blobSetTagsResponse.RawResponse.StatusCode, chk.Equals, 204)
//
//	blobGetTagsResponse, err := bbClient.GetTags(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(blobGetTagsResponse.RawResponse.StatusCode, chk.Equals, 200)
//	blobTagsSet := blobGetTagsResponse.Tags.BlobTagSet
//	c.Assert(blobTagsSet, chk.NotNil)
//	c.Assert(*blobTagsSet, chk.HasLen, 3)
//	for _, blobTag := range *blobTagsSet {
//		c.Assert(blobTagsMap[*blobTag.Key], chk.Equals, *blobTag.Value)
//	}
//}
//
//func (s *azblobTestSuite) TestSetBlobTagsWithVID() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	bbClient, _ := getBlockBlobClient(c, containerClient)
//	blobTagsMap := map[string]string{
//		"Go":         "CPlusPlus",
//		"Python":     "CSharp",
//		"Javascript": "Android",
//	}
//	blockBlobUploadResp, err := bbClient.Upload(ctx, bytes.NewReader([]byte("data")), nil)
//	_assert.Nil(err)
//	c.Assert(blockBlobUploadResp.RawResponse.StatusCode, chk.Equals, 201)
//	versionId1 := blockBlobUploadResp.VersionID
//
//	blockBlobUploadResp, err = bbClient.Upload(ctx, bytes.NewReader([]byte("updated_data")), nil)
//	_assert.Nil(err)
//	c.Assert(blockBlobUploadResp.RawResponse.StatusCode, chk.Equals, 201)
//	versionId2 := blockBlobUploadResp.VersionID
//
//	setTagsBlobOptions := SetTagsBlobOptions{
//		BlobTagsMap: &blobTagsMap,
//		VersionID:   versionId1,
//	}
//	blobSetTagsResponse, err := bbClient.SetTags(ctx, &setTagsBlobOptions)
//	_assert.Nil(err)
//	c.Assert(blobSetTagsResponse.RawResponse.StatusCode, chk.Equals, 204)
//
//	getTagsBlobOptions1 := GetTagsBlobOptions{
//		VersionID: versionId1,
//	}
//	blobGetTagsResponse, err := bbClient.GetTags(ctx, &getTagsBlobOptions1)
//	_assert.Nil(err)
//	c.Assert(blobGetTagsResponse.RawResponse.StatusCode, chk.Equals, 200)
//	c.Assert(blobGetTagsResponse.Tags.BlobTagSet, chk.NotNil)
//	c.Assert(*blobGetTagsResponse.Tags.BlobTagSet, chk.HasLen, 3)
//	for _, blobTag := range *blobGetTagsResponse.Tags.BlobTagSet {
//		c.Assert(blobTagsMap[*blobTag.Key], chk.Equals, *blobTag.Value)
//	}
//
//	getTagsBlobOptions2 := GetTagsBlobOptions{
//		VersionID: versionId2,
//	}
//	blobGetTagsResponse, err = bbClient.GetTags(ctx, &getTagsBlobOptions2)
//	_assert.Nil(err)
//	c.Assert(blobGetTagsResponse.RawResponse.StatusCode, chk.Equals, 200)
//	c.Assert(blobGetTagsResponse.Tags.BlobTagSet, chk.IsNil)
//}
//
//func (s *azblobTestSuite) TestUploadBlockBlobWithSpecialCharactersInTags() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	bbClient, _ := getBlockBlobClient(c, containerClient)
//	blobTagsMap := map[string]string{
//		"+-./:=_ ": "firsttag",
//		"tag2":     "+-./:=_",
//		"+-./:=_1": "+-./:=_",
//	}
//
//	uploadBlockBlobOptions := UploadBlockBlobOptions{
//		Metadata:        &basicMetadata,
//		BlobHTTPHeaders: &basicHeaders,
//		BlobTagsMap:     &blobTagsMap,
//	}
//	blockBlobUploadResp, err := bbClient.Upload(ctx, bytes.NewReader([]byte("data")), &uploadBlockBlobOptions)
//	_assert.Nil(err)
//	// TODO: Check for metadata and header
//	c.Assert(blockBlobUploadResp.RawResponse.StatusCode, chk.Equals, 201)
//
//	blobGetTagsResponse, err := bbClient.GetTags(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(blobGetTagsResponse.RawResponse.StatusCode, chk.Equals, 200)
//	c.Assert(*blobGetTagsResponse.Tags.BlobTagSet, chk.HasLen, 3)
//	for _, blobTag := range *blobGetTagsResponse.Tags.BlobTagSet {
//		c.Assert(blobTagsMap[*blobTag.Key], chk.Equals, *blobTag.Value)
//	}
//}
//
//func (s *azblobTestSuite) TestStageBlockWithTags() {
//	blockIDIntToBase64 := func(blockID int) string {
//		binaryBlockID := (&[4]byte{})[:]
//		binary.LittleEndian.PutUint32(binaryBlockID, uint32(blockID))
//		return base64.StdEncoding.EncodeToString(binaryBlockID)
//	}
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//
//	bbClient := containerClient.NewBlockBlobClient(generateBlobName())
//
//	data := []string{"Azure ", "Storage ", "Block ", "Blob."}
//	base64BlockIDs := make([]string, len(data))
//
//	for index, d := range data {
//		base64BlockIDs[index] = blockIDIntToBase64(index)
//		resp, err := bbClient.StageBlock(ctx, base64BlockIDs[index], strings.NewReader(d), nil)
//		_assert.Nil(err)
//		c.Assert(resp.RawResponse.StatusCode, chk.Equals, 201)
//		c.Assert(*resp.Version, chk.Not(chk.Equals), "")
//	}
//
//	blobTagsMap := map[string]string{
//		"azure":    "bbClient",
//		"bbClient": "sdk",
//		"sdk":      "go",
//	}
//
//	commitBlockListOptions := CommitBlockListOptions{
//		BlobTagsMap: &blobTagsMap,
//	}
//	commitResp, err := bbClient.CommitBlockList(ctx, base64BlockIDs, &commitBlockListOptions)
//	_assert.Nil(err)
//	c.Assert(commitResp.VersionID, chk.NotNil)
//	versionId := commitResp.VersionID
//
//	contentResp, err := bbClient.Download(ctx, nil)
//	_assert.Nil(err)
//	contentData, err := ioutil.ReadAll(contentResp.Body(RetryReaderOptions{}))
//	c.Assert(contentData, chk.DeepEquals, []uint8(strings.Join(data, "")))
//
//	getTagsBlobOptions := GetTagsBlobOptions{
//		VersionID: versionId,
//	}
//	blobGetTagsResp, err := bbClient.GetTags(ctx, &getTagsBlobOptions)
//	_assert.Nil(err)
//	c.Assert(blobGetTagsResp, chk.NotNil)
//	c.Assert(*blobGetTagsResp.Tags.BlobTagSet, chk.HasLen, 3)
//	for _, blobTag := range *blobGetTagsResp.Tags.BlobTagSet {
//		c.Assert(blobTagsMap[*blobTag.Key], chk.Equals, *blobTag.Value)
//	}
//
//	blobGetTagsResp, err = bbClient.GetTags(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(blobGetTagsResp, chk.NotNil)
//	c.Assert(*blobGetTagsResp.Tags.BlobTagSet, chk.HasLen, 3)
//	for _, blobTag := range *blobGetTagsResp.Tags.BlobTagSet {
//		c.Assert(blobTagsMap[*blobTag.Key], chk.Equals, *blobTag.Value)
//	}
//}
//
//func (s *azblobTestSuite) TestStageBlockFromURLWithTags() {
//	bsu := getServiceClient(nil)
//	credential, err := getGenericCredential("")
//	if err != nil {
//		c.Fatal("Invalid credential")
//	}
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//
//	contentSize := 4 * 1024 * 1024 // 4MB
//	r, sourceData := getRandomDataAndReader(contentSize)
//	ctx := ctx // Use default Background context
//	srcBlob := containerClient.NewBlockBlobClient("sourceBlob")
//	destBlob := containerClient.NewBlockBlobClient("destBlob")
//
//	blobTagsMap := map[string]string{
//		"Go":         "CPlusPlus",
//		"Python":     "CSharp",
//		"Javascript": "Android",
//	}
//
//	uploadBlockBlobOptions := UploadBlockBlobOptions{
//		BlobTagsMap: &blobTagsMap,
//	}
//	uploadSrcResp, err := srcBlob.Upload(ctx, r, &uploadBlockBlobOptions)
//	_assert.Nil(err)
//	c.Assert(uploadSrcResp.RawResponse.StatusCode, chk.Equals, 201)
//
//	// Get source blob url with SAS for StageFromURL.
//	srcBlobParts := NewBlobURLParts(srcBlob.URL())
//
//	srcBlobParts.SAS, err = BlobSASSignatureValues{
//		Protocol:      SASProtocolHTTPS,                     // Users MUST use HTTPS (not HTTP)
//		ExpiryTime:    time.Now().UTC().Add(48 * time.Hour), // 48-hours before expiration
//		ContainerName: srcBlobParts.ContainerName,
//		BlobName:      srcBlobParts.BlobName,
//		Permissions:   BlobSASPermissions{Read: true}.String(),
//	}.NewSASQueryParameters(credential)
//	if err != nil {
//		c.Fatal(err)
//	}
//
//	srcBlobURLWithSAS := srcBlobParts.URL()
//
//	blockID1 := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%6d", 0)))
//	blockID2 := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%6d", 1)))
//
//	offset1, count1 := int64(0), int64(contentSize/2)
//	options1 := StageBlockFromURLOptions{
//		Offset: &offset1,
//		Count:  &count1,
//	}
//	stageResp1, err := destBlob.StageBlockFromURL(ctx, blockID1, srcBlobURLWithSAS, 0, &options1)
//	_assert.Nil(err)
//	c.Assert(stageResp1.RawResponse.StatusCode, chk.Equals, 201)
//	c.Assert(*stageResp1.RequestID, chk.Not(chk.Equals), "")
//	c.Assert(*stageResp1.Version, chk.Not(chk.Equals), "")
//	c.Assert(stageResp1.Date, chk.NotNil)
//	c.Assert((*stageResp1.Date).IsZero(), chk.Equals, false)
//
//	offset2, count2 := int64(contentSize/2), int64(CountToEnd)
//	options2 := StageBlockFromURLOptions{
//		Offset: &offset2,
//		Count:  &count2,
//	}
//	stageResp2, err := destBlob.StageBlockFromURL(ctx, blockID2, srcBlobURLWithSAS, 0, &options2)
//	_assert.Nil(err)
//	c.Assert(stageResp2.RawResponse.StatusCode, chk.Equals, 201)
//	c.Assert(*stageResp2.RequestID, chk.Not(chk.Equals), "")
//	c.Assert(*stageResp2.Version, chk.Not(chk.Equals), "")
//	c.Assert(stageResp2.Date, chk.NotNil)
//	c.Assert((*stageResp2.Date).IsZero(), chk.Equals, false)
//
//	blockList, err := destBlob.GetBlockList(ctx, BlockListAll, nil)
//	_assert.Nil(err)
//	c.Assert(blockList.RawResponse.StatusCode, chk.Equals, 200)
//	c.Assert(blockList.BlockList.CommittedBlocks, chk.IsNil)
//	c.Assert(*blockList.BlockList.UncommittedBlocks, chk.HasLen, 2)
//
//	commitBlockListOptions := CommitBlockListOptions{
//		BlobTagsMap: &blobTagsMap,
//	}
//	listResp, err := destBlob.CommitBlockList(ctx, []string{blockID1, blockID2}, &commitBlockListOptions)
//	_assert.Nil(err)
//	c.Assert(listResp.RawResponse.StatusCode, chk.Equals, 201)
//	//versionId := listResp.VersionID()
//
//	blobGetTagsResp, err := destBlob.GetTags(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(*blobGetTagsResp.Tags.BlobTagSet, chk.HasLen, 3)
//	for _, blobTag := range *blobGetTagsResp.Tags.BlobTagSet {
//		c.Assert(blobTagsMap[*blobTag.Key], chk.Equals, *blobTag.Value)
//	}
//
//	downloadResp, err := destBlob.Download(ctx, nil)
//	_assert.Nil(err)
//	destData, err := ioutil.ReadAll(downloadResp.Body(RetryReaderOptions{}))
//	_assert.Nil(err)
//	c.Assert(destData, chk.DeepEquals, sourceData)
//}
//
//func (s *azblobTestSuite) TestCopyBlockBlobFromURLWithTags() {
//	bsu := getServiceClient(nil)
//	credential, err := getGenericCredential("")
//	if err != nil {
//		c.Fatal("Invalid credential")
//	}
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//
//	contentSize := 8 * 1024 // 1MB
//	r, sourceData := getRandomDataAndReader(contentSize)
//	sourceDataMD5Value := md5.Sum(sourceData)
//	srcBlob := containerClient.NewBlockBlobClient("srcBlob")
//	destBlob := containerClient.NewBlockBlobClient("destBlob")
//
//	blobTagsMap := map[string]string{
//		"Go":         "CPlusPlus",
//		"Python":     "CSharp",
//		"Javascript": "Android",
//	}
//
//	uploadBlockBlobOptions := UploadBlockBlobOptions{
//		BlobTagsMap: &blobTagsMap,
//	}
//	uploadSrcResp, err := srcBlob.Upload(ctx, r, &uploadBlockBlobOptions)
//	_assert.Nil(err)
//	c.Assert(uploadSrcResp.RawResponse.StatusCode, chk.Equals, 201)
//
//	// Get source blob url with SAS for StageFromURL.
//	srcBlobParts := NewBlobURLParts(srcBlob.URL())
//
//	srcBlobParts.SAS, err = BlobSASSignatureValues{
//		Protocol:      SASProtocolHTTPS,                     // Users MUST use HTTPS (not HTTP)
//		ExpiryTime:    time.Now().UTC().Add(48 * time.Hour), // 48-hours before expiration
//		ContainerName: srcBlobParts.ContainerName,
//		BlobName:      srcBlobParts.BlobName,
//		Permissions:   BlobSASPermissions{Read: true}.String(),
//	}.NewSASQueryParameters(credential)
//	if err != nil {
//		c.Fatal(err)
//	}
//
//	srcBlobURLWithSAS := srcBlobParts.URL()
//	sourceContentMD5 := sourceDataMD5Value[:]
//	copyBlockBlobFromURLOptions1 := CopyBlockBlobFromURLOptions{
//		BlobTagsMap:      &map[string]string{"foo": "bar"},
//		SourceContentMD5: &sourceContentMD5,
//	}
//	resp, err := destBlob.CopyFromURL(ctx, srcBlobURLWithSAS, &copyBlockBlobFromURLOptions1)
//	_assert.Nil(err)
//	c.Assert(resp.RawResponse.StatusCode, chk.Equals, 202)
//	c.Assert(*resp.ETag, chk.Not(chk.Equals), "")
//	c.Assert(*resp.RequestID, chk.Not(chk.Equals), "")
//	c.Assert(*resp.Version, chk.Not(chk.Equals), "")
//	c.Assert((*resp.Date).IsZero(), chk.Equals, false)
//	c.Assert(*resp.CopyID, chk.Not(chk.Equals), "")
//	c.Assert(*resp.ContentMD5, chk.DeepEquals, sourceDataMD5Value[:])
//	c.Assert(*resp.CopyStatus, chk.DeepEquals, "success")
//
//	downloadResp, err := destBlob.Download(ctx, nil)
//	_assert.Nil(err)
//	destData, err := ioutil.ReadAll(downloadResp.Body(RetryReaderOptions{}))
//	_assert.Nil(err)
//	c.Assert(destData, chk.DeepEquals, sourceData)
//	c.Assert(*downloadResp.TagCount, chk.Equals, int64(1))
//
//	_, badMD5 := getRandomDataAndReader(16)
//	copyBlockBlobFromURLOptions2 := CopyBlockBlobFromURLOptions{
//		BlobTagsMap:      &blobTagsMap,
//		SourceContentMD5: &badMD5,
//	}
//	_, err = destBlob.CopyFromURL(ctx, srcBlobURLWithSAS, &copyBlockBlobFromURLOptions2)
//	_assert.NotNil(err)
//
//	copyBlockBlobFromURLOptions3 := CopyBlockBlobFromURLOptions{
//		BlobTagsMap: &blobTagsMap,
//	}
//	resp, err = destBlob.CopyFromURL(ctx, srcBlobURLWithSAS, &copyBlockBlobFromURLOptions3)
//	_assert.Nil(err)
//	c.Assert(resp.RawResponse.StatusCode, chk.Equals, 202)
//	c.Assert(resp.RawResponse.Header.Get("x-ms-content-crc64"), chk.Not(chk.Equals), "")
//}
//
//func (s *azblobTestSuite) TestGetPropertiesReturnsTagsCount() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	bbClient, _ := getBlockBlobClient(c, containerClient)
//
//	uploadBlockBlobOptions := UploadBlockBlobOptions{
//		BlobTagsMap:     &basicBlobTagsMap,
//		BlobHTTPHeaders: &basicHeaders,
//		Metadata:        &basicMetadata,
//	}
//	blockBlobUploadResp, err := bbClient.Upload(ctx, bytes.NewReader([]byte("data")), &uploadBlockBlobOptions)
//	_assert.Nil(err)
//	c.Assert(blockBlobUploadResp.RawResponse.StatusCode, chk.Equals, 201)
//
//	getPropertiesResponse, err := bbClient.GetProperties(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(*getPropertiesResponse.TagCount, chk.Equals, int64(3))
//
//	downloadResp, err := bbClient.Download(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(downloadResp, chk.NotNil)
//	c.Assert(downloadResp.RawResponse.Header.Get("x-ms-tag-count"), chk.Equals, "3")
//}
//
//func (s *azblobTestSuite) TestSetBlobTagForSnapshot() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	bbClient, _ := createNewBlockBlob(c, containerClient)
//	blobTagsMap := map[string]string{
//		"Microsoft Azure": "Azure Storage",
//		"Storage+SDK":     "SDK/GO",
//		"GO ":             ".Net",
//	}
//	setTagsBlobOptions := SetTagsBlobOptions{
//		BlobTagsMap: &blobTagsMap,
//	}
//	_, err := bbClient.SetTags(ctx, &setTagsBlobOptions)
//	_assert.Nil(err)
//
//	resp, err := bbClient.CreateSnapshot(ctx, nil)
//	_assert.Nil(err)
//
//	snapshotURL := bbClient.WithSnapshot(*resp.Snapshot)
//	resp2, err := snapshotURL.GetProperties(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(*resp2.TagCount, chk.Equals, int64(3))
//}
//
//// TODO: Once new pacer is done.
////func (s *azblobTestSuite) TestListBlobReturnsTags() {
////	bsu := getServiceClient()
////	containerClient, _ := createNewContainer(c, bsu)
////	defer deleteContainer(containerClient)
////	blobClient, blobName := createNewBlockBlob(c, containerClient)
////	blobTagsMap := BlobTagsMap{
////		"+-./:=_ ": "firsttag",
////		"tag2":     "+-./:=_",
////		"+-./:=_1": "+-./:=_",
////	}
////	resp, err := blobClient.SetTags(ctx, nil, nil, nil, nil, nil, nil, blobTagsMap)
////	_assert.Nil(err)
////	c.Assert(resp.StatusCode(), chk.Equals, 204)
////
////	listBlobResp, err := containerClient.ListBlobsFlatSegment(ctx, Marker{}, ListBlobsSegmentOptions{Details: BlobListingDetails{Tags: true}})
////
////	_assert.Nil(err)
////	c.Assert(listBlobResp.Segment.BlobItems[0].Name, chk.Equals, blobName)
////	c.Assert(listBlobResp.Segment.BlobItems[0].BlobTags.BlobTagSet, chk.HasLen, 3)
////	for _, blobTag := range listBlobResp.Segment.BlobItems[0].BlobTags.BlobTagSet {
////		c.Assert(blobTagsMap[blobTag.Key], chk.Equals, blobTag.Value)
////	}
////}
////
////func (s *azblobTestSuite) TestFindBlobsByTags() {
////	bsu := getServiceClient()
////	containerClient1, _ := createNewContainer(c, bsu)
////	defer deleteContainer(containerClient1)
////	containerClient2, _ := createNewContainer(c, bsu)
////	defer deleteContainer(containerClient2)
////	containerClient3, _ := createNewContainer(c, bsu)
////	defer deleteContainer(containerClient3)
////
////	blobTagsMap1 := BlobTagsMap{
////		"tag2": "tagsecond",
////		"tag3": "tagthird",
////	}
////	blobTagsMap2 := BlobTagsMap{
////		"tag1": "firsttag",
////		"tag2": "secondtag",
////		"tag3": "thirdtag",
////	}
////	blobURL11, _ := getBlockBlobURL(c, containerClient1)
////	_, err := blobURL11.Upload(ctx, bytes.NewReader([]byte("random data")), BlobHTTPHeaders{}, basicMetadata, BlobAccessConditions{}, DefaultAccessTier, blobTagsMap1, ClientProvidedKeyOptions{})
////	_assert.Nil(err)
////	blobURL12, _ := getBlockBlobURL(c, containerClient1)
////	_, err = blobURL12.Upload(ctx, bytes.NewReader([]byte("another random data")), BlobHTTPHeaders{}, basicMetadata, BlobAccessConditions{}, DefaultAccessTier, blobTagsMap2, ClientProvidedKeyOptions{})
////	_assert.Nil(err)
////
////	blobURL21, _ := getBlockBlobURL(c, containerClient2)
////	_, err = blobURL21.Upload(ctx, bytes.NewReader([]byte("random data")), BlobHTTPHeaders{}, basicMetadata, BlobAccessConditions{}, DefaultAccessTier, nil, ClientProvidedKeyOptions{})
////	_assert.Nil(err)
////	blobURL22, _ := getBlockBlobURL(c, containerClient2)
////	_, err = blobURL22.Upload(ctx, bytes.NewReader([]byte("another random data")), BlobHTTPHeaders{}, basicMetadata, BlobAccessConditions{}, DefaultAccessTier, blobTagsMap2, ClientProvidedKeyOptions{})
////	_assert.Nil(err)
////
////	blobURL31, _ := getBlockBlobURL(c, containerClient3)
////	_, err = blobURL31.Upload(ctx, bytes.NewReader([]byte("random data")), BlobHTTPHeaders{}, basicMetadata, BlobAccessConditions{}, DefaultAccessTier, nil, ClientProvidedKeyOptions{})
////	_assert.Nil(err)
////
////	where := "\"tag4\"='fourthtag'"
////	lResp, err := bsu.FindBlobsByTags(ctx, nil, nil, &where, Marker{}, nil)
////	_assert.Nil(err)
////	c.Assert(lResp.Blobs, chk.HasLen, 0)
////
////	//where = "\"tag1\"='firsttag'AND\"tag2\"='secondtag'AND\"@container\"='"+ containerName1 + "'"
////	//TODO: Figure out how to do a composite query based on container.
////	where = "\"tag1\"='firsttag'AND\"tag2\"='secondtag'"
////
////	lResp, err = bsu.FindBlobsByTags(ctx, nil, nil, &where, Marker{}, nil)
////	_assert.Nil(err)
////
////	for _, blob := range lResp.Blobs {
////		c.Assert(blob.TagValue, chk.Equals, "firsttag")
////	}
////}
////
////func (s *azblobTestSuite) TestFilterBlobsUsingAccountSAS() {
////	accountName, accountKey := accountInfo()
////	credential, err := NewSharedKeyCredential(accountName, accountKey)
////	if err != nil {
////		c.Fail()
////	}
////
////	sasQueryParams, err := AccountSASSignatureValues{
////		Protocol:      SASProtocolHTTPS,
////		ExpiryTime:    time.Now().UTC().Add(48 * time.Hour),
////		Permissions:   AccountSASPermissions{Read: true, List: true, Write: true, DeletePreviousVersion: true, Tag: true, FilterByTags: true, Create: true}.String(),
////		Services:      AccountSASServices{Blob: true}.String(),
////		ResourceTypes: AccountSASResourceTypes{Service: true, Container: true, Object: true}.String(),
////	}.NewSASQueryParameters(credential)
////	if err != nil {
////		log.Fatal(err)
////	}
////
////	qp := sasQueryParams.Encode()
////	urlToSendToSomeone := fmt.Sprintf("https://%s.blob.core.windows.net?%s", accountName, qp)
////	u, _ := url.Parse(urlToSendToSomeone)
////	serviceURL := NewServiceURL(*u, NewPipeline(NewAnonymousCredential(), PipelineOptions{}))
////
////	containerName := generateContainerName()
////	containerClient := serviceURL.NewcontainerClient(containerName)
////	_, err = containerClient.Create(ctx, Metadata{}, PublicAccessNone)
////	defer containerClient.Delete(ctx, ContainerAccessConditions{})
////	if err != nil {
////		c.Fatal(err)
////	}
////
////	blobClient := containerClient.NewBlockBlobURL("temp")
////	_, err = blobClient.Upload(ctx, bytes.NewReader([]byte("random data")), BlobHTTPHeaders{}, basicMetadata, BlobAccessConditions{}, DefaultAccessTier, nil, ClientProvidedKeyOptions{})
////	if err != nil {
////		c.Fail()
////	}
////
////	blobTagsMap := BlobTagsMap{"tag1": "firsttag", "tag2": "secondtag", "tag3": "thirdtag"}
////	setBlobTagsResp, err := blobClient.SetTags(ctx, nil, nil, nil, nil, nil, nil, blobTagsMap)
////	_assert.Nil(err)
////	c.Assert(setBlobTagsResp.StatusCode(), chk.Equals, 204)
////
////	blobGetTagsResp, err := blobClient.GetTags(ctx, nil, nil, nil, nil, nil)
////	_assert.Nil(err)
////	c.Assert(blobGetTagsResp.StatusCode(), chk.Equals, 200)
////	c.Assert(blobGetTagsResp.BlobTagSet, chk.HasLen, 3)
////	for _, blobTag := range blobGetTagsResp.BlobTagSet {
////		c.Assert(blobTagsMap[blobTag.Key], chk.Equals, blobTag.Value)
////	}
////
////	time.Sleep(30 * time.Second)
////	where := "\"tag1\"='firsttag'AND\"tag2\"='secondtag'AND@container='" + containerName + "'"
////	_, err = serviceURL.FindBlobsByTags(ctx, nil, nil, &where, Marker{}, nil)
////	_assert.Nil(err)
////}
//
//func (s *azblobTestSuite) TestCreatePageBlobWithTags() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	contentSize := 1 * 1024
//	offset, count := int64(0), int64(contentSize)
//	uploadPagesOptions := UploadPagesOptions{
//		PageRange: &HttpRange{offset, count},
//	}
//	putResp, err := pbClient.UploadPages(ctx, getReaderToRandomBytes(1024), &uploadPagesOptions)
//	_assert.Nil(err)
//	c.Assert(putResp.RawResponse.StatusCode, chk.Equals, 201)
//	c.Assert(putResp.LastModified.IsZero(), chk.Equals, false)
//	c.Assert(putResp.ETag, chk.Not(chk.Equals), ETagNone)
//	c.Assert(putResp.Version, chk.Not(chk.Equals), "")
//
//	setTagsBlobOptions := SetTagsBlobOptions{
//		BlobTagsMap: &basicBlobTagsMap,
//	}
//	setTagResp, err := pbClient.SetTags(ctx, &setTagsBlobOptions)
//	_assert.Nil(err)
//	c.Assert(setTagResp.RawResponse.StatusCode, chk.Equals, 204)
//
//	gpResp, err := pbClient.GetProperties(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(gpResp, chk.NotNil)
//	c.Assert(*gpResp.TagCount, chk.Equals, int64(len(basicBlobTagsMap)))
//
//	blobGetTagsResponse, err := pbClient.GetTags(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(blobGetTagsResponse.RawResponse.StatusCode, chk.Equals, 200)
//	blobTagsSet := blobGetTagsResponse.Tags.BlobTagSet
//	c.Assert(blobTagsSet, chk.NotNil)
//	c.Assert(*blobTagsSet, chk.HasLen, len(basicBlobTagsMap))
//	for _, blobTag := range *blobTagsSet {
//		c.Assert(basicBlobTagsMap[*blobTag.Key], chk.Equals, *blobTag.Value)
//	}
//
//	modifiedBlobTags := map[string]string{
//		"a0z1u2r3e4": "b0l1o2b3",
//		"b0l1o2b3":   "s0d1k2",
//	}
//
//	setTagsBlobOptions2 := SetTagsBlobOptions{
//		BlobTagsMap: &modifiedBlobTags,
//	}
//	setTagResp, err = pbClient.SetTags(ctx, &setTagsBlobOptions2)
//	_assert.Nil(err)
//	c.Assert(setTagResp.RawResponse.StatusCode, chk.Equals, 204)
//
//	gpResp, err = pbClient.GetProperties(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(gpResp, chk.NotNil)
//	c.Assert(*gpResp.TagCount, chk.Equals, int64(len(modifiedBlobTags)))
//
//	blobGetTagsResponse, err = pbClient.GetTags(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(blobGetTagsResponse.RawResponse.StatusCode, chk.Equals, 200)
//	blobTagsSet = blobGetTagsResponse.Tags.BlobTagSet
//	c.Assert(blobTagsSet, chk.NotNil)
//	c.Assert(*blobTagsSet, chk.HasLen, len(modifiedBlobTags))
//	for _, blobTag := range *blobTagsSet {
//		c.Assert(modifiedBlobTags[*blobTag.Key], chk.Equals, *blobTag.Value)
//	}
//}
//
//func (s *azblobTestSuite) TestPageBlobSetBlobTagForSnapshot() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	pbClient, _ := createNewPageBlob(c, containerClient)
//
//	setTagsBlobOptions := SetTagsBlobOptions{
//		BlobTagsMap: &specialCharBlobTagsMap,
//	}
//	_, err := pbClient.SetTags(ctx, &setTagsBlobOptions)
//	_assert.Nil(err)
//
//	resp, err := pbClient.CreateSnapshot(ctx, nil)
//	_assert.Nil(err)
//
//	snapshotURL := pbClient.WithSnapshot(*resp.Snapshot)
//	resp2, err := snapshotURL.GetProperties(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(*resp2.TagCount, chk.Equals, int64(len(specialCharBlobTagsMap)))
//
//	blobGetTagsResponse, err := pbClient.GetTags(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(blobGetTagsResponse.RawResponse.StatusCode, chk.Equals, 200)
//	blobTagsSet := blobGetTagsResponse.Tags.BlobTagSet
//	c.Assert(blobTagsSet, chk.NotNil)
//	c.Assert(*blobTagsSet, chk.HasLen, len(specialCharBlobTagsMap))
//	for _, blobTag := range *blobTagsSet {
//		c.Assert(specialCharBlobTagsMap[*blobTag.Key], chk.Equals, *blobTag.Value)
//	}
//}
//
//func (s *azblobTestSuite) TestCreateAppendBlobWithTags() {
//	bsu := getServiceClient(nil)
//	containerClient, _ := createNewContainer(c, bsu)
//	defer deleteContainer(containerClient)
//	abClient, _ := getAppendBlobClient(c, containerClient)
//
//	createAppendBlobOptions := CreateAppendBlobOptions{
//		BlobTagsMap: &specialCharBlobTagsMap,
//	}
//	createResp, err := abClient.Create(ctx, &createAppendBlobOptions)
//	_assert.Nil(err)
//	c.Assert(createResp.VersionID, chk.NotNil)
//
//	_, err = abClient.GetProperties(ctx, nil)
//	_assert.Nil(err)
//
//	blobGetTagsResponse, err := abClient.GetTags(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(blobGetTagsResponse.RawResponse.StatusCode, chk.Equals, 200)
//	blobTagsSet := blobGetTagsResponse.Tags.BlobTagSet
//	c.Assert(blobTagsSet, chk.NotNil)
//	c.Assert(*blobTagsSet, chk.HasLen, len(specialCharBlobTagsMap))
//	for _, blobTag := range *blobTagsSet {
//		c.Assert(specialCharBlobTagsMap[*blobTag.Key], chk.Equals, *blobTag.Value)
//	}
//
//	resp, err := abClient.CreateSnapshot(ctx, nil)
//	_assert.Nil(err)
//
//	snapshotURL := abClient.WithSnapshot(*resp.Snapshot)
//	resp2, err := snapshotURL.GetProperties(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(*resp2.TagCount, chk.Equals, int64(len(specialCharBlobTagsMap)))
//
//	blobGetTagsResponse, err = abClient.GetTags(ctx, nil)
//	_assert.Nil(err)
//	c.Assert(blobGetTagsResponse.RawResponse.StatusCode, chk.Equals, 200)
//	blobTagsSet = blobGetTagsResponse.Tags.BlobTagSet
//	c.Assert(blobTagsSet, chk.NotNil)
//	c.Assert(*blobTagsSet, chk.HasLen, len(specialCharBlobTagsMap))
//	for _, blobTag := range *blobTagsSet {
//		c.Assert(specialCharBlobTagsMap[*blobTag.Key], chk.Equals, *blobTag.Value)
//	}
//}