defmodule HTTPBench do
  use Lazymaru.Router

  get do
    ~w({"Hello":"World!"}) |> text(200)
  end

  get "/:name" do
    %{ hello: params[:name] } |> json
  end

end
