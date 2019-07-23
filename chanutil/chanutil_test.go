package chanutil

import "testing"

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestMerge(t *testing.T) {
	t.Log("Testing Merge")
	{
		t.Log("\t Merges input from multiple channels into a single channel")
		{
		}

		t.Log("\t Closes output channel once all the input channels are closed")
		{
		}

		t.Log("\t Allows to close input and output channels using a shutdown channel")
		{
		}

		t.Log("\t Does not allow to write on the output channel\n")
		{
		}
	}
}

func TestSplit(t *testing.T) {
	t.Log("Testing Split")
	{
		t.Log("\t Broadcasts (duplicates) input from a single channel across multiple output channels")
		{
		}

		t.Log("\t Closes output channels once input channel has been closed and all values have been read from the output channels")
		{
		}

		t.Log("\t Allows to close input and output channels using a shutdown channel\n")
		{
		}
	}
}

func TestSplitRnd(t *testing.T) {
	t.Log("Testing SplitRnd")
	{
		t.Log("\t Distributes input from a single channel across multiple output channels randomly")
		{
		}

		t.Log("\t Closes output channels once input channel has been closed and values have been read from the output channels")
		{
		}

		t.Log("\t Allows to close input and output channels using a shutdown channel\n")
		{
		}
	}
}

func TestPipeline(t *testing.T) {
	t.Log("Testing Pipeline")
	{
		t.Log("\t Create an asynchronous pipeline which filters values from the input channel in order of the functions provided and publishes qualifying values to the output channel")
		{
		}

		t.Log("\t Shuts down the pipeline and the output channel once the input channel is closed and all the values have been read from the output channel")
		{
		}

		t.Log("\t Allows to close the input channel, the pipeline and output channel using a shutdown channel")
		{
		}

		t.Log("\t Does not allow to write on the output channel\n")
		{
		}
	}
}

func TestOrShutdown(t *testing.T) {
	t.Log("Testing OrShutdown")
	{
		t.Log("\t Closes output channel if any of the input channels are closed")
		{
		}

		t.Log("\t Does not allow to write on the output channel")
		{
		}

	}
}

func TestAndShutdown(t *testing.T) {
	t.Log("Testing AndShutdown")
	{
		t.Log("\t Closes output channel when all of the input channels are closed")
		{
		}
		t.Log("\t Does not allow to write on the output channel")
		{
		}
	}
}

func TestPool(t *testing.T) {
	t.Log("Testing Pool")
	{
		t.Log("\t Spawns 'poolSize' goroutines to process values read from the input channel")
		{
		}

		t.Log("\t Closes output channel when 'shutdown' channel is closed")
		{
		}

		t.Log("\t Exits all worker goroutines when 'shutdown' channel is invoked")
		{
		}

		t.Log("\t Closes output channel when input channel is closed")
		{
		}
	}
}
