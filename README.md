# qalam

https://www.devdungeon.com/content/working-files-go

Simple time based filewriter to log data to disk. Has buffer support to flush data from memory to disk to reduce disk IO


Usage:


		c := New(filepath.Join("/tmp/isco.%Y%m%d%H%M"))
		c.Write("isco is better than bale")



TODO:

		Add multiple test cases