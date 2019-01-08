class Bencoding

  attr_reader :data

  
  def initialize(data)
    @data = data

  end


  def parse

    subset = []
    number = ""


    length = @data.split("").length
    counter = 0

    while counter < length do 
      
      
      if @data[counter] == "i"
        while @data[couner+1] != "e" do
          number << @data[counter+1]
          counter+=1
        end
      end
      counter =+1
    end
    subset
  end

end


bencoding = Bencoding.new("i10e")
puts bencoding.parse
