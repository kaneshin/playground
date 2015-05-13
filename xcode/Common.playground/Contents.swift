//: Playground - noun: a place where people can play

import Cocoa

var str = "{\"hello\": \"world\"}"
var data = str.dataUsingEncoding(NSUTF8StringEncoding, allowLossyConversion: true)
var json = NSJSONSerialization.JSONObjectWithData(data!, options: .MutableContainers, error: nil) as! NSDictionary
