#!/usr/bin/env python
# -*- coding: utf-8 -*-

from notion.client import NotionClient
from notion.block import (
    BulletedListBlock,
    HeaderBlock,
    PageBlock,
    SubheaderBlock)
from notion.block import TextBlock
import os

token = os.environ.get('NOTION_TOKEN')
client = NotionClient(token_v2=token)

url = os.environ.get('NOTION_PARENT_URL')
page = client.get_block(url)

added_page = page.children.add_new(PageBlock, title="<PAGE_TITLE>")
child_page = client.get_block(added_page.id)
child_page.children.add_new(
    HeaderBlock, title="heading1")
child_page.children.add_new(
    TextBlock, title="foo bar")
child_page.children.add_new(
    SubheaderBlock, title="heading2")
child_page.children.add_new(
    BulletedListBlock, title="foo")
child_page.children.add_new(
    BulletedListBlock, title="bar")
